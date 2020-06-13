package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Piochat/GoTwit/db"
)

//GetBanner para traer el Banner del usuario
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID not Foud", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open(PathBanner + profile.Banner)
	if err != nil {
		http.Error(w, "User Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copying image", http.StatusBadRequest)
	}
}
