package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Piochat/GoTwit/db"
)

//ReadTweets funcion que encuentra los tweets a leer
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The id is required", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "The page is required", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "The page value is invalid", http.StatusBadRequest)
		return
	}

	result, status := db.ReadTweets(ID, int64(page))
	if !status {
		http.Error(w, "Error reading tweet", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
