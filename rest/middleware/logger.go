package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger middleware - প্রতিটি request log করবে
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// মূল handler এ বানানো
		next.ServeHTTP(w, r)

		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}



