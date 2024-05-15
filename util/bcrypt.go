package util

import "golang.org/x/crypto/bcrypt"

func GenerateBcryptHash(password string) (string, error) {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passBytes), nil
}
