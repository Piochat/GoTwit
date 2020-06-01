package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Piochat/GoTwit/db"

	"github.com/Piochat/GoTwit/models"
)

//ModProfile actualiza los datos del perfil
func ModProfile(w http.ResponseWriter, r *http.Request) {
	var (
		t      models.User
		status bool
	)

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Data Wrong "+err.Error(), 400)
		return
	}

	status, err = db.ModRe(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error modifying user "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "An attempt to modify the user failed ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
