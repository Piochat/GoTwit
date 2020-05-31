package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Piochat/GoTwit/db"

	"github.com/Piochat/GoTwit/models"
)

//Register Funcion que registra un usuario en la base de datos
func Register(w http.ResponseWriter, r *http.Request) {
	// Estructura usuario
	var t models.User

	// Request Body se destruye en memoria
	// Es un stream (Todo lo del body se carga en User)
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		print("Error al registrar", err.Error())
		http.Error(w, "[Error Register]: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "[Error Email Empty]: "+err.Error(), 400)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "[Error Password Wrong (I need 8 char min)]: "+err.Error(), 400)
		return
	}

	_, found, _ := db.CheckUserExist(t.Email)
	if found {
		http.Error(w, "[Error User Already Exist]", 400)
		return
	}

	_, status, err := db.InsertRe(t)
	if err != nil {
		http.Error(w, "[Error Sing Up Failed (DB)]: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Sing Up Field (Die)", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
