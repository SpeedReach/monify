package main

import (
	"github.com/joho/godotenv"
	"monify/lib/utils"
	"monify/services/media"
)

func main() {
	//Load secrets
	_ = godotenv.Load()
	secrets, err := utils.LoadSecrets(utils.LoadEnv())
	if err != nil {
		panic(err)
	}
	infra, err := media.Setup(media.NewConfig(secrets))
	if err != nil {
		panic(err)
	}
	media.NewServer(infra).Start()
}
