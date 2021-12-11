package test

import (
	"E-Culture-API/models"
	"testing"
)

func TestCreatePlace(t *testing.T) {
	err := models.CreatePlace("Giuseppe", "xxx", "Via boh, 89", "BlaBla", 1)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
