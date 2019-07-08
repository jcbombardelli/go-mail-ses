package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	log.Println("\n Init AWS Mail Service ")

	key := cfg.Section("AWS SES").Key("aws_key_id")
	pass := cfg.Section("AWS SES").Key("aws_secret_key")
	region := cfg.Section("AWS SES").Key("aws_region")
	SetConfiguration(key.Value(), pass.Value(), region.Value())

	email := Email{
		From:    "no-reply@globalcode.com.br",
		To:      "jc.bombardelli@live.com",
		Subject: "Mail",
		Text:    "Text",
		HTML:    "<h1>oo</h1>",
		ReplyTo: "no-reply@globalcode.com.br",
	}

	ret := SendEmail(email)
	fmt.Println(ret)
}
