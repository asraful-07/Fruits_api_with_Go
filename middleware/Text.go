package middleware

import "net/http"

type Middlewaree func(http.Handler) http.Handler

type Managers struct {
	globalMiddlewares []Middlewaree
}

func SoManager() *Managers {
	return &Managers{
		globalMiddlewares: make([]Middlewaree, 0),
	}
}

func (mngr *Managers) Use(middlewares  ...Middlewaree) *Managers {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return  mngr
}

func (mngr *Managers) with(handler http.Handler, middlwares ...Middlewaree) http.Handler{
	h := handler

	for i := len(middlwares) - 1; i >= 0; i-- {
		h = middlwares[i](h)
	}

	for i := len(mngr.globalMiddlewares) -1; i >= 0; i-- {
		h = mngr.globalMiddlewares[i](h)
	}
	return  h
}

func (mng *Manager) Wrap(handler http.Handler) http.Handler {
	h := handler 
	for i := len(mng.globalMiddlewares) - 1; i >= 0; i-- {
		h = mng.globalMiddlewares[i](h)
	}
	return h
}