package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Piochat/GoTwit/db"
)

//ReadTweetsFollower lee los tweets de nuestros followers
func ReadTweetsFollower(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "the page parameter should be sent", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "the page parameter must be sent as an integer greater than zero", http.StatusBadRequest)
		return
	}

	result, status := db.ReadTweetsFollowers(IDUsuario, page)
	if !status {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
