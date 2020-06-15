package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Piochat/GoTwit/db"

	"github.com/Piochat/GoTwit/models"
)

//ConsultRelation chequea si hay relación entre 2 usuarios
func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UserID = IDUsuario
	t.UserRelationID = ID

	var resp models.RespuestaConsultaRelacion
	status, err := db.ConsultRelation(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
