package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func ComparePassword(hashedPass string, payloadPass []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), payloadPass)

	return err == nil
}
