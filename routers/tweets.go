package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Piochat/GoTwit/db"

	"github.com/Piochat/GoTwit/models"
)

//SendTweet envia el tweet a la base de datos
func SendTweet(w http.ResponseWriter, r *http.Request) {
	var tw models.Tweet
	var status bool

	err := json.NewDecoder(r.Body).Decode(&tw)

	item := models.InfoTweet{
		UserID:  IDUsuario,
		Mensaje: tw.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err = db.InsertTweet(item)
	if err != nil {
		http.Error(w, "Failed to send the Tweet "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Failed -- Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
