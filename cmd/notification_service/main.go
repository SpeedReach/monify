package main

import (
	"github.com/joho/godotenv"
	"monify/lib/utils"
	"monify/services/notification"
)

func main() {
	//Load secrets
	_ = godotenv.Load()
	secrets, err := utils.LoadSecrets(utils.LoadEnv())
	if err != nil {
		panic(err)
	}
	notification.Start(notification.NewConfig(secrets))
}
