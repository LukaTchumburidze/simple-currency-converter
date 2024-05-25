package db

import (
	"fmt"
	"github.com/LukaTchumburidze/simple-currency-converter/cfg"
	"github.com/LukaTchumburidze/simple-currency-converter/entity"
	"github.com/LukaTchumburidze/simple-currency-converter/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Database)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Logger.Info("connected to database")

	err = DB.AutoMigrate(&entity.Request{})
	if err != nil {
		log.Logger.Error("failed on request entity creation")
		panic(err)
	}
	err = DB.AutoMigrate(&entity.Response{})
	if err != nil {
		log.Logger.Error("failed on response entity creation")
		panic(err)
	}
	log.Logger.Info("migration finished")
}
