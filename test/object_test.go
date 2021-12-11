package test

import (
	"E-Culture-API/models"
	"testing"
)

func TestCreateObject(t *testing.T) {
	err := models.CreateObject("Venere di Nike", "BlaBla", "0", "xxxx", 1)
	if err != nil {
		t.Errorf("error should be nil but is %v", err)
	}
}
