package product

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
	"github.com/teclegacy/golang-ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewProductHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/products", h.handleGetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handlePostProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handlePostProduct(w http.ResponseWriter, r *http.Request) {
	// Parse Payload
	var payload types.Product
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate Payload
	if err := utils.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Create Product
	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	p := types.Product{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
		CreatedAt:   createdAt,
	}

	// Query the db
	err := h.store.CreateProduct(p)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating product: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}
