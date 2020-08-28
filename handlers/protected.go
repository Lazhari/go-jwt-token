package handlers

import (
	"fmt"
	"net/http"

	"github.com/lazhari/web-jwt/middleware"
)

// ProtectedHandler Protect route
func (c Controller) ProtectedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID := middleware.GetUserID(ctx)
		fmt.Fprintln(w, "Protect", userID)
	}
}
