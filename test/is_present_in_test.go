package test

import (
	"E-Culture-API/models"
	"testing"
)

func TestAddObjectToPath(t *testing.T) {
	err := models.AddObjectToPath(1, 1, 1)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
