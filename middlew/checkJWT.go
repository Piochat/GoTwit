package middlew

import (
	"net/http"

	"github.com/Piochat/GoTwit/routers"
)

//ValidateJWT valida el JWT que nos llega por la petición
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error Token Failed! "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
