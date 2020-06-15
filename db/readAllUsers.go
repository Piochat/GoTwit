package db

import (
	"context"
	"log"
	"time"

	"github.com/Piochat/GoTwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadAllUsers lee lo usuarios registrados en el sistema.
func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("users")

	var results []*models.User
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	//Al usar Find retora un cursor
	cur, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, false
	}

	//Iteramos el cursor
	var found, include bool
	for cur.Next(ctx) {
		//Cargamos un modelo para contener los objetos del cursor
		//Decodificados
		var s models.User
		err = cur.Decode(&s)
		if err != nil {
			log.Println(err.Error(), "Cursor")
			return results, false
		}

		var r models.Relacion
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false
		found, err = ConsultRelation(r)
		if tipo == "new" && found == false {
			include = true
		}
		if tipo == "follow" && found == true {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		//Limpiar los datos del usuario, pertinentes al perfil
		//y la contraseña
		if include == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		log.Println(err.Error(), "Error cursor")
	}
	cur.Close(ctx)

	return results, true
}
