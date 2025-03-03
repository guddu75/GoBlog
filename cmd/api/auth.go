package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
	"github.com/guddu75/goblog/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

type UserWithToken struct {
	User  store.User `json:"user"`
	Token string     `json:"token"`
}

// registerUserHandler godoc
//
//	@Summary		Register a new user
//	@Description	Register a new user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User credentials"
//	@Success		201		{object}	userWithToken			"User registered"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	app.logger.Infow("Inside registerUserHandler", "method", r.Method, "path", r.URL.Path)

	var payload RegisterUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user := store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// hash the password
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	token := uuid.New().String()

	hash := sha256.Sum256([]byte(token))

	hashedToken := hex.EncodeToString(hash[:])

	err := app.store.Users.CreateAndInvite(ctx, &user, hashedToken, app.config.mail.exp)

	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestError(w, r, err)
		case store.ErrDuplicateUsername:
			app.badRequestError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	// TODO mail

	userWithToken := UserWithToken{
		User:  user,
		Token: token,
	}

	if err := writeJSON(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
