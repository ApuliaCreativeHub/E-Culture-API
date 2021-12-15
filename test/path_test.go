package test

import (
	"E-Culture-API/models"
	"testing"
)

func TestCreatePath(t *testing.T) {
	err := models.CreatePath("Path1", "BlaBla", 1)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}

func TestAddObjectToPath(t *testing.T) {
	err := models.AddObjectToPath(1, 1, 1)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
