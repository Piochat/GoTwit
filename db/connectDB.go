package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCon es una variable de conexión
var MongoCon = ConnectionDB()
var clietnsOptions = options.Client().ApplyURI("mongodb+srv://piocha:koke@cluster0-km1uz.mongodb.net/test?retryWrites=true&w=majority")

//ConnectionDB funcion
//Conexión a la basea de datos
func ConnectionDB() *mongo.Client {
	/*
	* Contexto: (No comparte cosas de la API en los llamados de esta)
	* Espacio en memoria, para acceder y compartir.
	* También se usan para controlar los cuelgues usando TimeOut
	* TODO Por defecto con todos los valores default y todo control
	 */
	client, err := mongo.Connect(context.TODO(), clietnsOptions)

	if err != nil {
		log.Fatalln("[ConnectionDB] Error", err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalln("[Ping]", err.Error())
		return client
	}

	log.Println("Connection Successful!!!!!")

	return client
}

//CheckConnetcion funcion
//Realiza un ping a la base de datos
func CheckConnetcion() bool {
	err := MongoCon.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true
}
