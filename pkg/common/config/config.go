package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port   string `mapstructure:"PORT" validate:"required"`
	DB_URL string `mapstructure:"DB_URL" validate:"required"`
}

var Env *Configuration = &Configuration{}

type ConfigModule struct {
}

func (module *ConfigModule) Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%v", err)
	}

	if err := viper.Unmarshal(Env); err != nil {
		log.Fatalf("%v", err)
	}

	validate := validator.New()

	if err := validate.Struct(Env); err != nil {
		log.Fatalf("%v", err)
	}
}
