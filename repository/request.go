package repository

import (
	"github.com/LukaTchumburidze/simple-currency-converter/db"
	"github.com/LukaTchumburidze/simple-currency-converter/entity"
)

type RequestRepository struct {
}

func (r RequestRepository) StoreRequest(request entity.Request) (entity.Request, error) {
	err := db.DB.Create(request).Find(&request).Error
	return request, err
}
