package helper

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}

	return string(passwordHash), nil
}

func ComparePassword(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	return err == nil

}
