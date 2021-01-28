package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet es la estructura con la que devolveremos los Tweets de un usuario determinado*/
type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
