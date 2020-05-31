package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Piochat/GoTwit/jwt"

	"github.com/Piochat/GoTwit/db"
	"github.com/Piochat/GoTwit/models"
)

//Login contorl para iniciar sesión
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	// Cargando datos de json a usuer t
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error User or Password "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "I need the email", 400)
		return
	}

	doc, exist := db.TryToLogin(t.Email, t.Password)
	if !exist {
		http.Error(w, "Error User or Password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "Error: Token Generate Failed"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Save a cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
