package middleware

import "net/http"

// Middleware টাইপ
type Middleware func(http.Handler) http.Handler

// Manager struct (global middleware store করবে)
type Manager struct {
	globalMiddlewares []Middleware
}

// NewManager - নতুন Manager বানানো
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// Use - Global middleware যোগ করা
func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

// With - Local + Global middleware চেইন করে final handler return
func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	// Local middleware (reverse এ wrap)
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	// Global middleware (reverse এ wrap)
	for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
		h = mngr.globalMiddlewares[i](h)
	}

	return h
}


/*
// --- Local middleware wrap (reverse with range) ---
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	// --- Global middleware wrap (reverse with range) ---
	for _, globalMiddlewares := range mngr.globalMiddlewares {
		h = globalMiddlewares(h)
	}
*/