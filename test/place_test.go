package test

import (
	"E-Culture-API/models"
	"gorm.io/gorm/utils"
	"testing"
)

func TestReadByUserId(t *testing.T) {
	place := models.Place{UserID: 1}
	places, err := place.ReadByUserId()
	if err != nil {
		return
	}
	utils.AssertEqual([]models.Place{{ID: 1, Name: "Place1", PhotoPath: "xxx", Address: "Via boh, 89", Description: "BlaBla", UserID: 1}}, places)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
