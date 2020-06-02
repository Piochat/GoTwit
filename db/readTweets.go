package db

import (
	"context"
	"log"
	"time"

	"github.com/Piochat/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadTweets Devulve una lista y su paginacion.
func ReadTweets(ID string, page int64) ([]*models.GetTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("tweet")

	var result []*models.GetTweet

	condition := bson.M{
		"userid": ID,
	}

	// Control para poder paginar la cantidad tweets
	option := options.Find()
	option.SetLimit(20)
	option.SetSort(bson.D{{Key: "fecha", Value: -1}})
	option.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, option)
	if err != nil {
		log.Fatalln(err.Error())
		return result, false
	}

	//Cargando los registros válidos en el resultado
	for cursor.Next(context.TODO()) {
		var reg models.GetTweet
		err := cursor.Decode(&reg)

		if err != nil {
			return result, false
		}
		result = append(result, &reg)
	}

	return result, true
}
