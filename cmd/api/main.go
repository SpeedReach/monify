package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"monify/lib/utils"
	"monify/services/api"
	"monify/services/api/infra"
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
	serverCfg := api.NewConfigFromSecrets(secrets)

	//initialize server
	s := api.NewServer(serverCfg, resources)

	port := "2302"
	//start listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	println("Server is running on port " + port)
	err = s.Start(lis)
	if err != nil {
		log.Fatal(err)
	}
}
