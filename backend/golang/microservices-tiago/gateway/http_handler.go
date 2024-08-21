package main

import (
	"net/http"

	"github.com/teclegacy/ms/oms/common"
	pb "github.com/teclegacy/ms/oms/common/api"
)

type Handler struct {
	//gateway client instance of order service
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *Handler {
	return &Handler{client}
}

func (s *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/v1/customers/{customerId}/orders", s.HandlerCreateOrder)
}
func (s *Handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")

	var items []*pb.ItemsWithQuantity

	if err := common.ReadJSON(r, &items); err != nil {
		common.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Call the grpc server to create order in the order service
	s.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerId,
		Items:      items,
	})

}
