package db

import (
	"context"
	"time"

	"github.com/Piochat/GoTwit/models"
)

//InsertRel guardad el follow de los usuarios
func InsertRel(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("relationship")

	_, err := collection.InsertOne(ctx, t)

	return (err == nil), err
}
