package repository

import (
	"github.com/LukaTchumburidze/simple-currency-converter/db"
	"github.com/LukaTchumburidze/simple-currency-converter/entity"
)

type ResponseRepository struct {
}

func (r ResponseRepository) StoreResponse(response entity.Response) (entity.Response, error) {
	err := db.DB.Create(&response).Find(&response).Error
	return response, err
}
