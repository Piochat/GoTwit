package routers

import (
	"net/http"

	"github.com/Piochat/GoTwit/db"
	"github.com/Piochat/GoTwit/models"
)

//BajaRelacion realiza el borrado del follow entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UserID = IDUsuario
	t.UserRelationID = ID

	status, err := db.DeleteRelationship(t)
	if err != nil {
		http.Error(w, "An error occurred while deleting relationship "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Error deleting follow", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
