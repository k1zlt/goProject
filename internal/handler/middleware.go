package handler

import (
	"context"
	"first/helper"
	"first/logs"
	"net/http"
)

// authorizeUser is a middleware function that authorizes the user based on the provided token in the request header.
// It validates the token, extracts the user ID, and sets it in the request context for further processing.
//
// Parameters:
// - next: The next http.Handler in the middleware chain.
//
// Returns:
// - An http.Handler that performs the user authorization and invokes the next handler.
func (h *Handler) authorizeUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			logs.Log(logs.Error, "empty Authentication header", r.Method, r.URL.Path, 0)
			helper.ErrorResponse(w, "empty Authentication header", http.StatusUnauthorized, 1)
			return
		}
		userID, err := h.services.ParseToken(tokenString)
		if err != nil {
			logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, userID)
			helper.ErrorResponse(w, err.Error(), http.StatusUnauthorized, 1)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// checkEndpointPermission is a middleware function that checks if the user identified by the "userID" value in the request context
// has permission to access the requested endpoint. It queries the user's permissions for the given endpoint path and allows or denies access accordingly.
//
// Parameters:
// - next: The next http.Handler in the middleware chain.
//
// Returns:
// - An http.Handler that performs endpoint permission checks and invokes the next handler.
func (h *Handler) checkEndpointPermission(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(int)

		hasPermission, err := h.services.GetUserPermissionForEndpoint(userID, r.URL.Path)
		if err != nil {
			logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, userID)
			helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
			return
		}

		if !hasPermission {
			logs.Log(logs.Error, "rejected in accessing the path", r.Method, r.URL.Path, userID)
			helper.ErrorResponse(w, "rejected in accessing the path", http.StatusUnauthorized, 1)
			return
		}

		next.ServeHTTP(w, r)
	})
}
