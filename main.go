package main

import (
	utils "emailSender/util"
	"log"
)

func main() {

	viperConfig, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config, err=> ", err)
	}
	address := viperConfig.EmailSenderAddress
	println("email address =>>", address)
}
