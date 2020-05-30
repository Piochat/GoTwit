package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Piochat/GoTwit/models"
)

//InsertRe Inserta el usuario en la base de datos
func InsertRe(user models.User) (string, bool, error) {
	//Contexto para evitar el cuelgue de la operación
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	// Se ejecuata al final de toda ejecusion pese a las exceptions
	// Cancel -- Cancela el contexto
	defer cancel()

	// Etnrar a la base de datos y la colección
	database := MongoCon.Database("gotwit")
	collection := database.Collection("users")

	// Encriptar contraseña
	user.Password, _ = EncryptPassword(user.Password)

	// Inserta el usuario, pero valida el contexto
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	// Convierta el id de un ObjectID de mongo
	// a un dato de go
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
