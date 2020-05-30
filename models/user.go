package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User una estructura para el registro de usuario
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"frechaNacimiento" json:"nomfrechaNacimientobre,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" josn:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" josn:"sitioWeb,omitempty"`
}
