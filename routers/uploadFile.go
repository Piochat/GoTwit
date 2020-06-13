package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Piochat/GoTwit/db"

	"github.com/Piochat/GoTwit/models"
)

//PathAvatar ruta de las imagenes
const PathAvatar string = "uploads/avatar/"

//UploadAvatar funion para mandar un avatar
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = PathAvatar + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error Image upload fail "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error Image copy fail "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = archivo[len(PathAvatar):]
	status, err = db.ModRe(user, IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error Image save fail "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
