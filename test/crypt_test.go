package test

import (
	"E-Culture-API/utils"
	"fmt"
	"testing"
)

func TestSetConfigShouldBeValid(t *testing.T) {
	err := utils.SetConfigPath("conf/conf.yml")
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}

func TestNewJSONWebTokenShouldBeValid(t *testing.T) {
	err := utils.SetConfigPath("../conf/conf.yml")
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
	token, err := utils.NewJSONWebToken()
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
	fmt.Println(token)
}
