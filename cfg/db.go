package cfg

import "os"

var DB DBConfig

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func init() {
	DB = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	}
}
