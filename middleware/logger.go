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

		log.Printf("[Logger] %s %s", r.Method, r.URL.Path)

		// মূল handler এ পাঠানো
		next.ServeHTTP(w, r)

		log.Printf("[Logger] Completed in %v", time.Since(start))
	})
}



