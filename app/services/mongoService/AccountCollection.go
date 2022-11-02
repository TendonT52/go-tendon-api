package mongoService

// import (
// 	"context"
// 	"time"

// 	"github.com/spf13/viper"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // var AccountCollection = initAccountCollection()

// type InsertUserToDB interface {
// 	InsertUserToDB() interface{}
// }

// func initAccountCollection() *mongo.Collection{
// 	return 
// }

// func AddUser(firstName, lastName, email, hashPassword string) (string, error) {
// 	type userDocument struct {
// 		FirstName    string             `bson:"firstName"`
// 		LastName     string             `bson:"lastName"`
// 		Email        string             `bson:"email"`
// 		HashPassword string             `bson:"password"`
// 		Role         string             `bson:"role"`
// 		updateAt     primitive.DateTime `bson:"updateAt"`
// 	}
// 	user := userDocument{
// 		FirstName:    firstName,
// 		LastName:     lastName,
// 		Email:        email,
// 		HashPassword: hashPassword,
// 		Role:         Student,
// 		updateAt:     primitive.NewDateTimeFromTime(time.Now()),
// 	}
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		viper.GetDuration("mongo.collection.account.addUserTimeOut"),
// 	)
// 	defer cancel()
// 	insertOne, err := accountCollection.InsertOne(ctx, user)
// 	if err != nil {
// 		return "", err
// 	}
// 	return insertOne.InsertedID.(primitive.ObjectID).Hex(), nil
// }

// func GetUserById(id string) (*models.User, error) {
// 	objId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	filter := bson.D{{Key: "_id", Value: objId}}
// 	var user models.User
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		viper.GetDuration("mongo.collection.account.getUserTimeout"),
// 	)
// 	defer cancel()
// 	err = accountCollection.FindOne(ctx, filter).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (account *AccountCollection) GetUserByEmail(email string) (*models.User, error) {
// 	filter := bson.D{{Key: "email", Value: email}}
// 	var user models.User
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		viper.GetDuration("mongo.collection.account.getUserTimeout"),
// 	)
// 	defer cancel()
// 	err := account.collection.FindOne(ctx, filter).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (account *AccountCollection) DeleteUserById(id string) (int, error) {
// 	objId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	filter := bson.D{{Key: "_id", Value: objId}}
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		viper.GetDuration("mongo.collection.account.getUserTimeOut"),
// 	)
// 	defer cancel()
// 	res, err := account.collection.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return int(res.DeletedCount), nil
// }
