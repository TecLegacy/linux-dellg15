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

	t.Run("fails when user provides invalid payload", func(t *testing.T) {

		//Request to /register POST
		// Invalid Payload
		payload := types.RegisterUserPayload{
			FirstName: "keshav",
			LastName:  "kumar",
			Email:     "INVALID-DATA",
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

			t.Errorf("Expected status code %v, got %v", http.StatusOK, rr.Code)
			t.Logf("Response body: %s", rr.Body.String())
		}

	})
	t.Run("fails when user Exists", func(t *testing.T) {

		//Request to /register POST
		// Invalid Payload
		payload := types.RegisterUserPayload{
			FirstName: "keshav",
			LastName:  "kumar",
			Email:     "DATA@gmail.com",
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

			t.Errorf("Expected status code %v, got %v", http.StatusOK, rr.Code)
			t.Logf("Response body: %s", rr.Body.String())
		}

	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	// Mock user NOT exists
	// return nil, fmt.Errorf("user not found")

	// User Found
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
