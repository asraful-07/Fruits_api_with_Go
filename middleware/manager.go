package middleware

import "net/http"

// Middleware টাইপ
type Middleware func(http.Handler) http.Handler

// Manager struct (ইচ্ছে করলে config বা অন্য জিনিস রাখতে পারবে)
type Manager struct{}

// NewManager - নতুন Manager বানানোর জন্য
func NewManager() *Manager {
	return &Manager{}
}

// With - একাধিক middleware একসাথে চেইন করবে
func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		h := handler

		// পেছন দিক থেকে middleware লাগানো হচ্ছে
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}

		return h
	}
}
