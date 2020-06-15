package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Piochat/GoTwit/db"
)

//ListUsers consulta todos los usuarios y los pagina
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	numPage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		http.Error(w, "You must send the page parameter as an integer greater than zero", http.StatusBadRequest)
		return
	}

	result, status := db.ReadAllUsers(IDUsuario, numPage, search, typeUser)
	if !status {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
