package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim es la estructura usada para procesar el JWT */
type Claim struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Email string             `json:"email"`
	jwt.StandardClaims
}
