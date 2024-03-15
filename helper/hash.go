package helper

import "golang.org/x/crypto/bcrypt"

func Encrypt(s string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(p, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(s))
	return err == nil
}