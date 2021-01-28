package models

/*TweetMensaje captura del Body, el mensaje que nos llega*/
type TweetMensaje struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
	//Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
