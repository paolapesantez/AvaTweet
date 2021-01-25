package models

import "time"

/*TweetUser es el formato que tendrá nuestro Tweet en la BD*/
type TweetUser struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
