package auth // Change this to the name of your package

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/teclegacy/golang-ecom/config"
)

/**
The token is successfully created and not an empty string.
The signing method is as expected (HMAC in this case).
The userID claim matches the input.
The expiration (expiredAt) is within the expected range.
**/

func TestCreateJWT(t *testing.T) {
	// Example secret key for testing
	secret := []byte("testSecretKey")
	userID := 123

	// Call the function to generate a JWT
	tokenString, err := CreateJWT(secret, userID)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Ensure the token is not empty
	if tokenString == "" {
		t.Fatalf("Expected a token to be generated, got an empty string")
	}

	// Parse the token to check if it's valid and check claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Fatalf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Convert userID to string because it's stored as a string in the token
		if strconv.Itoa(userID) != claims["userID"] {
			t.Errorf("Expected userID %d, got %s", userID, claims["userID"])
		}

		// Check expiration is roughly correct, allowing a small margin for test execution time
		if exp, ok := claims["expiredAt"].(float64); !ok || time.Until(time.Unix(int64(exp), 0)) > time.Second*time.Duration(config.Envs.JWTExpiration) {
			t.Errorf("Token expiration is not within the expected range")
		}
	} else {
		t.Fatalf("Failed to parse token: %v", err)
	}
}
