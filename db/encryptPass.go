package db

import (
	"golang.org/x/crypto/bcrypt"
)

//EncryptPassword función para encriptar contraseñas
func EncryptPassword(pass string) (string, error) {
	// 2^costo : Cantidad de veces que encripta el password
	// Mayor cantidad de procesamiento mayor demora
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	return string(bytes), err
}
