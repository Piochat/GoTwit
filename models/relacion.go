package models

//Relacion modelo para grabar el follow de los usuarios
type Relacion struct {
	UserID         string `bson:"userid" json:"userid"`
	UserRelationID string `bson:"userRelationid" json:"userReationid"`
}
