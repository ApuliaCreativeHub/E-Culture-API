package controllers

import (
	"E-Culture-API/utils"
	"net/smtp"
	"strings"
)

var configPath = "conf/conf.yml"

func SendMail(msg string, recivers []string) error {
	err := utils.SetConfigPath(configPath)
	if err != nil {
		return err
	}
	conf, err := utils.GetConfig()
	if err != nil {
		return err
	}

	from := conf.Email.Address
	pass := conf.Email.Password

	var to string
	if len(recivers) == 1 {
		to = recivers[len(recivers)-1]
	} else {
		to = strings.Join(recivers, ",")
	}

	msg = "From: " + from + "\n" +
		"To: " + to + "\n" + msg

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
