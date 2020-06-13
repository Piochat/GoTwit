package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Piochat/GoTwit/db"

	"github.com/Piochat/GoTwit/models"
)

//PathBanner ruta de las imagenes
const PathBanner string = "uploads/banner/"

//UploadBanner funion para mandar un Banner
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = PathBanner + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error Banner upload fail "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error Banner copy fail "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = archivo[len(PathBanner):]
	status, err = db.ModRe(user, IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error Banner save fail "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
