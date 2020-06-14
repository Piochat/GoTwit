package routers

import (
	"net/http"

	"github.com/Piochat/GoTwit/db"
	"github.com/Piochat/GoTwit/models"
)

//AltaRealacion realiza el registro del follow entre usuarios
func AltaRealacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Erro ID not found", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UserID = IDUsuario
	t.UserRelationID = ID

	status, err := db.InsertRel(t)
	if err != nil {
		http.Error(w, "An error occurred while inserting relationship "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Error inserting follow", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
