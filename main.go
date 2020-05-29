package main

import (
	"log"

	"github.com/Piochat/GoTwit/handlers"

	"github.com/Piochat/GoTwit/db"
)

func main() {
	if !db.CheckConnetcion() {
		log.Fatalln("Error en la conexión de la base de datos")
		return
	}

	handlers.Controllers()
}
