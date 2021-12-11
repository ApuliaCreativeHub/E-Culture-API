package test

import (
	"E-Culture-API/models"
	"testing"
)

func TestCreateZone(t *testing.T) {
	err := models.CreateZone("Zona1", "BlaBla", 1)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
