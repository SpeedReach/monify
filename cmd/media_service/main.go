package main

import (
	"github.com/joho/godotenv"
	"monify/lib/utils"
	"monify/services/media"
)

func main() {
	//Load secrets
	_ = godotenv.Load()
	env := utils.LoadEnv()
	secrets, err := utils.LoadSecrets(env)
	if err != nil {
		panic(err)
	}
	infra, err := media.Setup(media.NewConfig(env, secrets))
	if err != nil {
		panic(err)
	}
	media.NewServer(infra).Start()
}
