package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TweetsFollowers para consultar
type TweetsFollowers struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"userid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
