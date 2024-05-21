package config

import (
	"basecode/model"
	"log"
)

func (config *Config) AutoMigration() {
	err := DB.AutoMigrate(model.User{}, model.Product{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("success migration")
}
