package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Piochat/GoTwit/models"
)

//ModRe funcion para modificar registro de los usuarios
func ModRe(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoCon.Database("gotwit")
	collection := database.Collection("users")

	registro := make(map[string]interface{})

	if len(user.Nombre) > 0 {
		registro["nombre"] = user.Nombre
	}
	if len(user.Apellidos) > 0 {
		registro["apellidos"] = user.Apellidos
	}
	registro["fechaNacimiento"] = user.FechaNacimiento
	if len(user.Avatar) > 0 {
		registro["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registro["banner"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		registro["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		registro["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		registro["sitioWeb"] = user.SitioWeb
	}

	updString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := collection.UpdateOne(ctx, filtro, updString)
	if err != nil {
		return false, err
	}

	return true, nil
}
