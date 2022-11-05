package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	password := []byte(p)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func ComparePassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
