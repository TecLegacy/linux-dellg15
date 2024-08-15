package main

import (
	"net/http"
)

type Handler struct {
	//gateway
}

func NewHandler() *Handler {
	return &Handler{}
}

func (s *Handler) RegisterRoutes(router *http.ServeMux) {

}
