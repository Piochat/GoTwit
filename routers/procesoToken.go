package routers

import (
	"errors"
	"strings"

	"github.com/Piochat/GoTwit/db"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Piochat/GoTwit/models"
)

//Email para consultar
var Email string

//IDUsuario identificador
var IDUsuario string

//ProcesoToken Ayuda a extraer los valores el token a lo largo del programa
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Tumadre")
	claims := &models.Claim{}

	// Separemos el token en Array
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("format toke invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	// Validación propia del Token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	}) // Verificamos si los valores son válidos
	if err == nil {
		_, found, _ := db.CheckUserExist(claims.Email)
		if found {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, found, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalid")
	}

	return claims, true, tk, err
}
