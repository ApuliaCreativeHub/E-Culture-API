package test

import (
	"E-Culture-API/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	err := models.CreateUser("Giuseppe", "Fortunato", "giuseppe.fortunato@pianto.it", "IoPiango1!", false)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
