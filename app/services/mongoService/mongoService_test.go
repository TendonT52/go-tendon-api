package mongoService

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"
)


func TestMain(m *testing.M) {
	code := m.Run()
	clean()
	DisconnectMongo()
	os.Exit(code)
}

func loadConfig(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../..")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("fatal error config file: %s", err)
	}
	log.Println("Load config success")
	connectionString := viper.GetString("mongo.connection")
	log.Printf("Connection String, %v \n", connectionString)
}

func clean() {
	res, err := AccountCollection.collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("Error while delete document %v", err)
		return
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

func shutdown() {
	DisconnectMongo()
}

func TestPingMongo(t *testing.T) {
	PingMongo()
}