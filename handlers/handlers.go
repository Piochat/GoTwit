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
	router.HandleFunc("/login", middlew.CheckDataBase(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDataBase(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modprofile", middlew.CheckDataBase(middlew.ValidateJWT(routers.ModProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDataBase(middlew.ValidateJWT(routers.SendTweet))).Methods("POST")
	router.HandleFunc("/readtweet", middlew.CheckDataBase(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deltweet", middlew.CheckDataBase(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDataBase(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadBanner", middlew.CheckDataBase(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDataBase(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/getBanner", middlew.CheckDataBase(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/followVo", middlew.CheckDataBase(middlew.ValidateJWT(routers.AltaRealacion))).Methods("POST")
	router.HandleFunc("/delFollow", middlew.CheckDataBase(middlew.ValidateJWT(routers.BajaRelacion))).Methods("DELETE")
	// End Rutas

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8484"
	}

	log.Println("[PORT:]", PORT)
	// CORS
	// Permiso a cualquiera
	handler := cors.AllowAll().Handler(router)
	log.Fatalln(http.ListenAndServe(":"+PORT, handler))
}
