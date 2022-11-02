package config

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

func init() {
	log.Println("Loading config...")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../..")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("Load config success")
}