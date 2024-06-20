package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
)

func TestUserRegisterHandler(t *testing.T) {
	//Handler -> mockUserStore
	mockUserStore := new(mockUserStore)
	handler := NewHandler(mockUserStore)

	t.Run("it fails when user payload is invalid", func(t *testing.T) {

		//Request to /register POST
		// Payload
		payload := types.RegisterUserPayload{
			FirstName: "keshav",
			LastName:  "kumar",
			Email:     "invalid",
			Password:  "asd",
		}

		marshalledPayload, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/register", handler.handleRegister)

		r.ServeHTTP(rr, req)

		// Assert response status code
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
		}

	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
