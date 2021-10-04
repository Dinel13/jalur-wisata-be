package main

import (
	"net/http"
)

// enablr Cors for all routes
func (a *application) enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// if r.Method == "OPTIONS" {
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
