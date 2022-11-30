package config

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port   string `mapstructure:"PORT" validate:"required"`
	DB_URL string `mapstructure:"DB_URL" validate:"required"`
}

var Value *Configuration = &Configuration{}

func LoadConfig() {
	// TODO: Create Log That read what module is being load
	fmt.Println("Loading Config")
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("env FILE NOT FOUND")
		log.Fatalf("%v", err)
	}

	if err := viper.Unmarshal(Value); err != nil {
		fmt.Println("Unable to decode into struct")
		log.Fatalf("%v", err)
	}

	validate := validator.New()
	fmt.Println("Validate Config")

	if err := validate.Struct(Value); err != nil {
		fmt.Println("Configuration validation failed")
		log.Fatalf("%v", err)
	}
}
