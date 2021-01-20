package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario contiene todos los campos que se almacenaran para cada usuario de avatweet
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty`
	Email           string             `bson:"email" json:"email,omitempty`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty`
	Password        string             `bson:"password" json:"password,omitempty`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty`
	Banner          string             `bson:"banner" json:"banner,omitempty`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty`
}
