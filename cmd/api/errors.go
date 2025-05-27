package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {

	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path)

	writeJSONError(w, http.StatusForbidden, "forbidden")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnf("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorw("conflict request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) unAuthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnf("Unauthorized request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusUnauthorized, "UnAuthorized request "+err.Error())
}

func (app *application) unAuthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnf("Basic Unauthorized request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted" charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, "UnAuthorized request")
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {

	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded , retry after "+retryAfter)
}
