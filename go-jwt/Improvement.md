### Tests for register user

```
package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/teclegacy/golang-ecom/types"
	"github.com/teclegacy/golang-ecom/utils"
	"github.com/stretchr/testify/mock"
	"github.com/go-playground/validator/v10"
)

func TestUserRegisterHandler(t *testing.T) {
	mockUserStore := new(MockUserStore)
	handler := NewHandler(mockUserStore)

	t.Run("it fails when user payload is invalid", func(t *testing.T) {
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

		// Mock the validator to return an error
		utils.Validator = validator.New()
		err = utils.Validator.Struct(payload)
		if err == nil {
			t.Fatal("expected validation error but got none")
		}

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v but got %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("it fails when user already exists", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password123",
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

		mockUserStore.On("GetUserByEmail", payload.Email).Return(&types.User{}, nil)

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v but got %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("it fails when password hashing fails", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password123",
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

		mockUserStore.On("GetUserByEmail", payload.Email).Return(nil, errors.New("not found"))

		// Mock the password hashing function to return an error
		utils.HashPassword = func(password string) (string, error) {
			return "", errors.New("hashing error")
		}

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status %v but got %v", http.StatusInternalServerError, rr.Code)
		}
	})

	t.Run("it registers a user successfully", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password123",
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

		mockUserStore.On("GetUserByEmail", payload.Email).Return(nil, errors.New("not found"))

		utils.HashPassword = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		mockUserStore.On("CreateUser", mock.Anything).Return(nil)

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status %v but got %v", http.StatusCreated, rr.Code)
		}
	})
}

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*types.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserStore) GetUserByID(id int) (*types.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*types.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserStore) CreateUser(user types.User) error {
	return m.Called(user).Error(0)
}

```
