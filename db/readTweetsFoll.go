package db

import (
	"context"
	"time"

	"github.com/Piochat/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadTweetsFollowers Lee los tweets de mis seguidores
func ReadTweetsFollowers(ID string, page int) ([]models.TweetsFollowers, bool) {
	var result []models.TweetsFollowers

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("relationship")

	skip := (page - 1) * 20

	//Union de tablas usando mongodb
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	//Unwind procesa la información para ser operable
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	//Aggregata framework para la union de colecciones en mongo
	cur, err := collection.Aggregate(ctx, conditions)
	err = cur.All(ctx, &result)

	return result, (err == nil)
}
