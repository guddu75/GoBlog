package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/guddu75/goblog/internal/store"
)

type userKey string

const userCtx userKey = "user"

// GetUser lists all existing accounts
//
//	@Summary		Fetches a user profile
//	@Description	Fetches a user profile by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"User ID"
//	@Success		200		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/users/{userId} [get]
//	@Security		ApiKeyAuth
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {

	user := getUserFromCtx(r)
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

// Follow User
//
//	@Summary		Follow a user
//	@Description	Follow a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int		true	"User ID"
//	@Success		204		{string}	string	"User unfollowed"
//	@Failure		400		{object}	error	"User payload missing"
//	@Failure		404		{object}	error	"User not found"
//	@Router			/users/{userId}/follow [put]
//	@Security		ApiKeyAuth
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromCtx(r)

	var payload FollowUser

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	// TODO : Revert back after authenication is implemented
	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Unfollow User
//
//	@Summary		Unfollow a user
//	@Description	Unfollow a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int		true	"User ID"
//	@Success		204		{string}	string	"User unfollowed"
//	@Failure		400		{object}	error	"User payload missing"
//	@Failure		404		{object}	error	"User not found"
//	@Router			/users/{userId}/unfollow [put]
//	@Security		ApiKeyAuth
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromCtx(r)

	var payload FollowUser

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	// TODO : Revert back after authenication is implemented
	if err := app.store.Followers.Unfollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Activate User
//
//	@Summary		Activates/Registers a user
//	@Description	Activates/Registers a user by token
//	@Tags			users
//	@Produce		json
//	@Param			token	path		string	true	"Inviatation token"
//	@Success		204		{string}	string	"User activated"
//	@Failure		500		{object}	error
//	@Failure		404		{object}	error
//	@Router			/users/activate/{token} [put]
//	@Security		ApiKeyAuth
func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	ctx := r.Context()

	err := app.store.Users.Activate(ctx, token)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.badRequestError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, ""); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")

		userID, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			app.badRequestError(w, r, err)
			return
		}
		ctx := r.Context()

		user, err := app.store.Users.GetByID(ctx, userID)

		if errors.Is(err, store.ErrNotFound) {
			app.notFoundError(w, r, err)
			return
		}

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
