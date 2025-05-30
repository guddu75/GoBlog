package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/guddu75/goblog/internal/store"
)

// func (app *application) LoggerMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		timeStart := time.Now()
// 		app.logger.Infow("request", "method", r.Method, "path", r.URL.Path , "time" , )
// 		next.ServeHTTP(w, r)

// 	})
// }

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Inside AuthTokenMiddleware")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unAuthorizedErrorResponse(w, r, fmt.Errorf("header is missing"))
			return
		}
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unAuthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}

		token := parts[1]

		jwtToken, err := app.auth.ValidateToken(token)
		if err != nil {
			app.unAuthorizedErrorResponse(w, r, fmt.Errorf("invalid token: %v", err))
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)

		fmt.Println("Printing userID here", userID)

		if err != nil {
			app.unAuthorizedErrorResponse(w, r, fmt.Errorf("invalid user ID in token: %v", err))
			return
		}
		ctx := r.Context()
		// user, err := app.store.Users.GetByID(ctx, userID)
		// if err != nil {
		// 	app.unAuthorizedErrorResponse(w, r, err)
		// 	return
		// }

		user, err := app.getUser(ctx, userID)

		if err != nil {
			app.unAuthorizedErrorResponse(w, r, fmt.Errorf("failed to get user: %v", err))
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				app.unAuthorizedBasicErrorResponse(w, r, fmt.Errorf("unauthorized request"))
				return
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Basic" {
				app.unAuthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}

			decoded, err := base64.StdEncoding.DecodeString(parts[1])

			if err != nil {
				app.unAuthorizedBasicErrorResponse(w, r, err)
				return
			}

			username := app.config.auth.basic.username
			password := app.config.auth.basic.password

			creds := strings.SplitN(string(decoded), ":", 2)

			if len(creds) != 2 || creds[0] != username || creds[1] != password {
				app.unAuthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}

func (app *application) checkPostOwnership(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Infow("Inside checkPostOwnership")
		user := getUserFromCtx(r)
		post := getPostFromCtx(r)

		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		allowed, err := app.checkRolePrecedence(r.Context(), user, requiredRole)

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		if !allowed {
			app.forbiddenResponse(w, r)
			return
		}

		app.logger.Infow("Exiting from checkRolePrecedence ", "allowed", allowed)

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkRolePrecedence(ctx context.Context, user *store.User, roleName string) (bool, error) {
	role, err := app.store.Roles.GetByName(ctx, roleName)

	fmt.Println("Printing role here")

	fmt.Println(role)

	if err != nil {
		return false, err
	}

	return user.Role.Level >= role.Level, nil
}

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {

	if !app.config.redisCfg.enabled {
		return app.store.Users.GetByID(ctx, userID)
	}

	user, err := app.cacheStorage.Users.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		app.logger.Infow("User not found in cache", "userID", userID)
		user, err = app.store.Users.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}
		if err := app.cacheStorage.Users.Set(ctx, user); err != nil {
			return nil, err
		}
	}

	return user, nil

}

func (app *application) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// app.logger.Infow("Inside RateLimiterMiddleware")
		if app.config.rateLimiter.Enabled {
			// app.logger.Infow("remote address", "remote address", r.RemoteAddr)
			if allow, retryAfter := app.rateLimiter.Allow(r.RemoteAddr); !allow {
				app.rateLimitExceededResponse(w, r, retryAfter.String())
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
