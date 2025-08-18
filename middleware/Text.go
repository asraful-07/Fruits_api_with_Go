package middleware

// import "net/http"

// type Middleware func(http.Handler) http.Handler

// type Manager struct {
// 	globalMiddlewares []Middleware
// }

// func NewManager() *Manager{
// 	return &Manager{
// 		globalMiddlewares: make([]Middleware, 0),
// 	}
// }

// func (mngr *Manager) Use(Middlewares ...Middleware) *Manager{
// 	mngr.globalMiddlewares = append(mngr.globalMiddlewares, Middlewares...)
// 	return mngr
// }

// func (mngr *Manager) With(handler http.Handler,Middlewares ...Middleware)  http.Handler{
// 	h := handler

// 	for i := len(Middlewares)  - 1; i >= 0; i--  {
//      h = Middlewares[i](h)
// 	}

// 	for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
// 		h = mngr.globalMiddlewares[i](h)
// 	}
// 	return  h
// }