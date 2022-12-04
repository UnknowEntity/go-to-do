package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"web/todo/pkg/common/config"
)

var DB *gorm.DB

type DBModule struct {
}

func (module *DBModule) Init() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Env.DB_URL), &gorm.Config{})

	if err != nil {
		log.Fatalf("%v", err)
	}
}
