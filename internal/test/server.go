package test

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"monify/internal"
	"monify/internal/utils"
	monify "monify/protobuf"
	"net"
	"sync"
	"testing"
)

var initialized = false
var lock = sync.Mutex{}
var client Client

type Client struct {
	monify.AuthServiceClient
	monify.GroupServiceClient
}

func startUnitTestServer(t *testing.T, lis net.Listener) {
	//Load secrets
	_ = godotenv.Load()
	secrets, err := utils.LoadSecrets("dev")
	if err != nil {
		panic(err)
	}

	//setup infrastructure
	resources := SetupTestResource(t)
	// load server config
	serverCfg := internal.NewConfigFromSecrets(secrets)

	//initialize server
	s := internal.NewServer(serverCfg, resources)

	//start listening
	go func() {
		err := s.Start(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func GetTestClient(t *testing.T) Client {
	lock.Lock()
	if initialized {
		return client
	}
	initialized = true

	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)
	startUnitTestServer(t, lis)

	conn, err := grpc.Dial("",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	client = Client{
		AuthServiceClient:  monify.NewAuthServiceClient(conn),
		GroupServiceClient: monify.NewGroupServiceClient(conn),
	}
	lock.Unlock()
	return client
}
