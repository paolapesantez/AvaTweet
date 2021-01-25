package models

import "time"

/*Tweet es el formato que tendrá nuestro Tweet en la BD*/
type Tweet struct {
	Mensaje string    `bson:"mensaje" json:"mensaje"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
