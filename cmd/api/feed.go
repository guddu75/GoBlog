package main

import (
	"net/http"

	"github.com/guddu75/goblog/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request context.
	//userID := r.Context().Value("userID").(string)

	fq := store.PaginatedFeedQuery{
		Limit:  10,
		Offset: 0,
		Sort:   "desc",
		Tags:   []string{},
		Search: "",
		Since:  "",
		Until:  "",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	// Parse the query string parameters.

	ctx := r.Context()

	// Fetch the user feed from the store.
	feed, err := app.store.Posts.GetUserFeed(ctx, int64(315), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Write the feed to the response.
	err = app.jsonResponse(w, http.StatusOK, feed)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}
