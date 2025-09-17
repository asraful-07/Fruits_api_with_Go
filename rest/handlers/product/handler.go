package product

import (
	"fruits-api/repo"
	"fruits-api/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	fruitsRepo  repo.FruitsRepo 
}

func NewHandler(middlewares *middleware.Middlewares, fruitsRepo  repo.FruitsRepo ) *Handler {
	return &Handler{
		middlewares: middlewares,
		fruitsRepo: fruitsRepo,
	}
}