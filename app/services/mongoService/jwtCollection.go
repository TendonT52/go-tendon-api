package mongoService

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type JwtCollection struct {
	collection *mongo.Collection
}

func NewJwtCollection(client *mongo.Client) *JwtCollection {
	return &JwtCollection{
		collection: client.Database(viper.GetString("mongo.name")).Collection(viper.GetString("mongo.jwtCollection.name")),
	}
}

// func (jwt *JwtCollection) AddToken(token models.AccessToken) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()
// 	insertOne, err := jwt.collection.InsertOne(ctx, token)
// 	if err != nil {
// 		return "", err
// 	}
// 	return insertOne.InsertedID.(primitive.ObjectID).Hex(), nil
// }
