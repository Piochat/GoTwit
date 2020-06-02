package routers

import (
	"net/http"

	"github.com/Piochat/GoTwit/db"
)

//DeleteTweet elimina el tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The id is required", http.StatusBadRequest)
		return
	}

	err := db.DelTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Error deleting the Tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
