package jwt

import (
	"time"

	"github.com/Piochat/GoTwit/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generando tokens para seguridad
func GenerateJWT(u models.User) (string, error) {
	// Clave para la firma del token
	miClave := []byte("Tumadre")

	payload := jwt.MapClaims{
		"email":            u.Email,
		"nombre":           u.Nombre,
		"appellidos":       u.Apellidos,
		"fecha_nacimiento": u.FechaNacimiento,
		"biografia":        u.Biografia,
		"ubicación":        u.Ubicacion,
		"sitioweb":         u.SitioWeb,
		"_id":              u.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	// Nuevo tokne con privilegios y signed
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
