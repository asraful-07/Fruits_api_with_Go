package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseWriterWrapper যাতে status code track করা যায়
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger middleware - প্রতিটি request log করবে, status code ও IP সহ
func Logger_Text(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// ResponseWriter wrap করা
		rw := &responseWriterWrapper{w, http.StatusOK}

		// Middleware শুরুতে log
		log.Printf("[Logger] Started %s %s | IP: %s", r.Method, r.URL.Path, r.RemoteAddr)

		// মূল handler এ পাঠানো
		next.ServeHTTP(rw, r)

		// Middleware শেষে log
		log.Printf("[Logger] Completed %s %s | Status: %d | Duration: %v | IP: %s",
			r.Method, r.URL.Path, rw.statusCode, time.Since(start), r.RemoteAddr)
	})
}
