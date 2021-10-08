// Package main
// Date:  2021/10/8 21:05
// Desc:
package main

import "net/http"

// Set Header X-Frame-Options: deny X-XSS-Protection: 1; mode=block
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}
