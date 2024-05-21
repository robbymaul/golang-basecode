package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Dialect  string
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

func (config *Config) GetConnection() error {
	var dsn string
	var dialector gorm.Dialector

	switch config.Dialect {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.User, config.Password, config.Host, config.Port, config.DBName)
		dialector = mysql.Open(dsn)
	case "postgresql":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host, config.User, config.Password, config.DBName, config.Port)
		dialector = postgres.Open(dsn)
	case "sqlserver":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			config.User, config.Password, config.Host, config.Port, config.DBName)
		dialector = sqlserver.Open(dsn)
	default:
		return fmt.Errorf("unsupported database dialect: %s", config.Dialect)
	}

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}
