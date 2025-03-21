package main

import (
	"expvar"
	"runtime"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/guddu75/goblog/internal/auth"
	"github.com/guddu75/goblog/internal/db"
	"github.com/guddu75/goblog/internal/env"
	"github.com/guddu75/goblog/internal/mailer"
	"github.com/guddu75/goblog/internal/ratelimiter"
	"github.com/guddu75/goblog/internal/store"
	"github.com/guddu75/goblog/internal/store/cache"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			GoBlog
//	@description	This is a sample server Blog server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

func main() {

	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:     env.GetString("REDIS_ADDR", "localhost:6379"),
			password: env.GetString("REDIS_PASSWORD", ""),
			db:       env.GetInt("REDIS_DB", 0),
			enabled:  env.GetBool("REDIS_ENABLED", true),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3,
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
			mailTrap: mailTrapConfig{
				apiKey:   env.GetString("MAILTRAP_API_KEY", ""),
				userName: env.GetString("MAILTRAP_USERNAME", ""),
				host:     env.GetString("MAILTRAP_HOST", ""),
				port:     env.GetInt("MAILTRAP_PORT", 1),
			},
		},
		auth: authConfig{
			basic: basicAuthConfig{
				username: env.GetString("BASIC_AUTH_USERNAME", "admin"),
				password: env.GetString("BASIC_AUTH_PASSWORD", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 7,
				iss:    "GoBlog",
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATE_LIMITER_REQUESTS_PER_TIME_FRAME", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}

	// Logger

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Info("Database connection pool established")

	// Cache

	var rdb *redis.Client

	if cfg.redisCfg.enabled {
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.redisCfg.addr,
			Password: cfg.redisCfg.password,
			DB:       cfg.redisCfg.db,
		})
		logger.Info("Redis connection established")
	}

	store := store.NewStorage(db)
	cacheStorage := cache.NewRedisStorage(rdb)

	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	mailer, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.host,
		cfg.mail.mailTrap.userName, cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail, cfg.mail.mailTrap.port)

	if err != nil {
		logger.Info("Error while creating mailer instance error -> ", err.Error())
	}

	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	app := &application{
		config:       cfg,
		store:        store,
		cacheStorage: cacheStorage,
		logger:       logger,
		mailer:       &mailer,
		auth:         jwtAuthenticator,
		rateLimiter:  rateLimiter,
	}

	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()

	// logger.Fatal(app.run(mux))

	app.logger.Info(app.run(mux))
}
