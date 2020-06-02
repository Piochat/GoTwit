package db

import (
	"context"
	"time"

	"github.com/Piochat/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertTweet funcion para creat tweets en la base de datos
func InsertTweet(t models.InfoTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("tweet")

	items := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := collection.InsertOne(ctx, items)
	if err != nil {
		return "", false, err
	}

	// Idicamos que del json que retorna al haber insetado extraiga el ObjectID
	// del ultimo campo insertado.
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
