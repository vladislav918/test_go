package handler

import (
	"net/http"
	repository "test-project/internal/database/postgres"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	productHandler := NewProductHandler(h.repository.Product)
	measureHandler := NewMeasureHandler(h.repository.Measure)

	mux.Handle("/product/", productHandler)
	mux.Handle("/measure/", measureHandler)

	return mux
}
