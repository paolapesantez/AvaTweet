package bd

import (
	"context"
	"time"

	"github.com/paolapesantez/avatweet/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarTweet graba el tweet en la BD*/
func InsertarTweet(tweet models.TweetUser) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("tweets")

	/*registro := bson.M{
		"userid":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}*/
	result, err := col.InsertOne(ctx, tweet)

	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
