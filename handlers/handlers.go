package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Piochat/GoTwit/middlew"
	"github.com/Piochat/GoTwit/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Controllers funcion para el control de las rutas
func Controllers() {
	// Contol para capturar el http y manejar request y response
	// Verfica el Body y el Header
	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/register", middlew.CheckDataBase(routers.Register)).Methods("POST")
	// End Rutas

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8484"
	}

	// CORS
	// Permiso a cualquiera
	handler := cors.AllowAll().Handler(router)
	log.Fatalln(http.ListenAndServe(":"+PORT, handler))
}
