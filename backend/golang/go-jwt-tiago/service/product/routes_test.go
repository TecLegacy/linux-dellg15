package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "time"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
)

func TestProductsHandler(t *testing.T) {
	mockStore := new(mockProductStore)
	productHandler := NewProductHandler(mockStore)

	t.Run("fails when user provides invalid payload ", func(t *testing.T) {

		// payload := types.Product{
		// 	ID:          1,
		// 	Name:        "Laptop",
		// 	Description: "A high-performance laptop.",
		// 	Image:       "https://example.com/laptop.jpg",
		// 	Price:       999.99,
		// 	Quantity:    10,
		// 	CreatedAt:   time.Now(),
		// }

		px := struct {
			ID string
		}{
			ID: "s",
		}

		p, err := json.Marshal(px)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(p))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		r := mux.NewRouter()
		r.HandleFunc("/products", productHandler.handlePostProduct)

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetAllProducts() ([]types.Product, error) {
	return nil, nil
}
func (m *mockProductStore) CreateProduct(product types.Product) error {
	return nil
}
func (s *mockProductStore) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	return nil, nil
}
