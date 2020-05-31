package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/Piochat/GoTwit/models"
)

//TryToLogin realiza un chequedo de login en la db
func TryToLogin(email string, password string) (models.User, bool) {
	user, found, id := CheckUserExist(email)
	if !found {
		log.Fatalln("[User not found]", id)
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	// Hash and Slice
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		log.Fatalln("[Password Wrong]", id)
		return user, false
	}
	return user, found
}
