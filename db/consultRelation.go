package db

import (
	"context"
	"log"
	"time"

	"github.com/Piochat/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultRelation consulta la relación entre 2 usuarios
func ConsultRelation(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("relationship")

	condition := bson.M{
		"userid":         t.UserID,
		"userRelationid": t.UserRelationID,
	}

	var result models.Relacion
	log.Println(result)
	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		log.Println(err.Error())
	}

	return (err == nil), err
}
