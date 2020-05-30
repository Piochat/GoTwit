package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Piochat/GoTwit/models"
)

//CheckUserExist función
//Comprueba si el usuario existe en la base de datos
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Etnrar a la base de datos y la colección
	databse := MongoCon.Database("gotwit")
	collection := databse.Collection("users")

	condicion := bson.M{"email": email}

	var result models.User

	// Lee en la base de datos
	err := collection.FindOne(ctx, condicion).Decode(&result)
	// Retorna el id como un string
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
