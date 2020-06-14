package db

import (
	"context"
	"time"

	"github.com/Piochat/GoTwit/models"
)

//DeleteRelationship borra la relación de la db
func DeleteRelationship(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("relationship")

	_, err := collection.DeleteOne(ctx, t)

	return (err == nil), err
}
