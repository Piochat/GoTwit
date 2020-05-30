package middlew

import (
	"net/http"

	"github.com/Piochat/GoTwit/db"
)

//CheckDataBase verfica la comunicación con la base de datos
func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnetcion() {
			http.Error(w, "Conexion perdida con la base de datos", 500)

			return
		}

		next.ServeHTTP(w, r)
	}
}
