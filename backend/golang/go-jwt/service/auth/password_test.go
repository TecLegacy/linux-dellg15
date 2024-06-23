package auth

import "testing"

func TestPassword(t *testing.T) {
	t.Helper()

	// Generate a hash password
	hashPass, err := HashPassword("somepass")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check if the hash password is not empty
	if string(hashPass) == "" {
		t.Errorf("Expected a non-empty hash string, but got an empty string")
	}

}

func TestHashPassword(t *testing.T) {
	t.Helper()

	// Generate a hash password
	hashPass, err := HashPassword("password123")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check if the hash password is not empty
	if string(hashPass) == "" {
		t.Errorf("Expected a non-empty hash string, but got an empty string")
	}

	// Compare the hash password with the original password
	val := ComparePassword(string(hashPass), []byte("password123"))
	if val == false {
		t.Errorf("Expected true, but got false")
	}
}
