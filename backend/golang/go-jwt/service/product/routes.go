package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
)

type Handler struct {
	store types.ProductStore
}

func NewProductHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {}
