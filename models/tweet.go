package models

//Tweet el contenido dentro del tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
