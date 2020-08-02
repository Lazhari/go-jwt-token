package controllers

import (
	"fmt"
	"net/http"
)

// ProtectedHandler Protect route
func (c Controller) ProtectedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Protect")
	}
}
