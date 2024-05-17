package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"monify/internal"
	"monify/internal/infra"
	"monify/internal/utils"
	"net"
)

func main() {
	//Load secrets
	_ = godotenv.Load()
	secrets, err := utils.LoadSecrets(utils.LoadEnv())
	if err != nil {
		panic(err)
	}

	//setup infrastructure
	infraCnf := infra.NewConfigFromSecrets(secrets)
	resources := infra.SetupResources(infraCnf)
	// load server config
	serverCfg := internal.NewConfigFromSecrets(secrets)

	//initialize server
	s := internal.NewServer(serverCfg, resources)

	//start listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
	if err != nil {
		panic(err)
	}
	err = s.Start(lis)
	if err != nil {
		log.Fatal(err)
	}
}
