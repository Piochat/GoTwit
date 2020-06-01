package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Piochat/GoTwit/db"
)

//ViewProfile funcion para conectar al perfil
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("_id")
	if len(ID) < 1 {
		http.Error(w, "ID don't sent", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error: Don't found Profile "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
