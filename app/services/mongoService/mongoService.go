package mongoService

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/TendonT52/go-tendon-api/app/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var	db *mongo.Database
var AccountCollection *mongo.Collection

type Insert interface{
	InsertMongo() (*mongo.Collection, interface{})
}

func init(){
	connectMongo()
	newDb()
	newAccountCollection()
}

func connectMongo(){
	log.Println("Connecting to mongo...")
	connectionString := viper.GetString("mongo.connection")
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("Connection String, %v \n", connectionString)
		log.Fatalf("Error while connect to mongo, %v", err)
	}
	log.Println("Connect mongo completed")
}

func newDb(){
	db = client.Database(viper.GetString("mongo.name"))
}

func newAccountCollection(){
	AccountCollection = db.Collection(viper.GetString("mongo.collection.account.name"))
}

func DisconnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Error while disconnect mongo, %v", err)
	}
	fmt.Println("Disconnect mongo completed")
}

func PingMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error while disconnect mongo, %v", err)
	}
	fmt.Println("Ping mongo Complete")
}

func InsertDoc(e interface{}) error {
	ins, ok := e.(Insert)
	if !ok {
		return errors.New("not implement insert interface")
	}
	col, doc := ins.InsertMongo()
	ctx, cancel := context.WithTimeout(
		context.Background(),
		viper.GetDuration("mongo.collection.account.insertTimeOut"),
	)
	defer cancel()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}