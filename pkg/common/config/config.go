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
	env *Configuration
}

func (module *ConfigModule) Init() {
	module.env = &Configuration{}
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%v", err)
	}

	if err := viper.Unmarshal(module.env); err != nil {
		log.Fatalf("%v", err)
	}

	validate := validator.New()

	if err := validate.Struct(module.env); err != nil {
		log.Fatalf("%v", err)
	}
}
