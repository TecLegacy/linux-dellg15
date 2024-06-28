package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
	"github.com/teclegacy/golang-ecom/utils"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewCartHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/cart/checkout", h.handleCartCheckout)
}

func (h *Handler) handleCartCheckout(w http.ResponseWriter, r *http.Request) {

	// Parse Cart item
	var cart types.CartCheckOutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate Payload
	if err := utils.Validator.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

}
