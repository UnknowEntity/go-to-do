package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port string `mapstructure:"PORT"`
}

var Value *Configuration = &Configuration{}

func LoadConfig() {
	fmt.Println("Loading Config")
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("env FILE NOT FOUND")
		log.Fatalf("%v", err)
	}

	err = viper.Unmarshal(Value)

	fmt.Printf("%v", Value)

	if err != nil {
		fmt.Println("Unable to decode into struct")
		log.Fatalf("%v", err)
	}
}
