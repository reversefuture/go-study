package db

import (
	"moduleDemo/shopping/models"
)

func LoadItem2(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
