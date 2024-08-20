package main

import (
	pb "github.com/teclegacy/ms/oms/common/api"
	"net/http"
)

type handler struct {
	//gateway instance
}

func NewHandler() *handler {
	return &handler{}
}

func (s *handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/v1/customers/{customerId}/orders", s.CreateOrder)
}
func (s *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	customerId := r.PathValue("customerId")

	var items []*pb.ItemsWithQuantity

}
