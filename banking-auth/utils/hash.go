package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *AppError) {
	pwd := toBytes(password)
	crypt, err := hashAndSalt(pwd)
	if err != nil {
		return "", err
	}
	return crypt, nil
}

func toBytes(pwd string) []byte {
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) (string, *AppError) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		LogError(err.Error())
		return "", NewUnexpectedError(err.Error())
	}
	return string(hash), nil
}

func ComparePasswords(hashedPwd string, plainPassword string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := toBytes(plainPassword)

	appError := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if appError != nil {
		LogError(appError.Error())
		return false
	}
	return true
}
