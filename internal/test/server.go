package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"monify/internal"
	"monify/internal/utils"
	monify "monify/protobuf"
	"net"
	"sync"
	"testing"
)

type State struct {
	initialized bool
	lock        sync.Mutex
	client      Client
	server      internal.Server
}

type Client struct {
	monify.AuthServiceClient
	monify.GroupServiceClient
	users       *map[string]string
	currentUser *string
}

var state State = State{
	initialized: false,
	lock:        sync.Mutex{},
}

func startUnitTestServer(t *testing.T, lis net.Listener) internal.Server {
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
	return s
}

func GetTestClient(t *testing.T) Client {
	state.lock.Lock()
	defer state.lock.Unlock()
	if state.initialized {
		return state.client
	}
	lis := createBuffListener()
	startUnitTestServer(t, lis)
	state.client = createClient(lis)
	state.initialized = true
	return state.client
}

func createBuffListener() *bufconn.Listener {
	return bufconn.Listen(101024 * 1024)
}

func createClient(lis *bufconn.Listener) Client {
	users := map[string]string{}
	currentUser := ""

	conn, err := grpc.Dial("",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			if currentUser != "" {
				md := metadata.Pairs("authorization", "Bearer "+users[currentUser])
				ctx = metadata.NewOutgoingContext(context.Background(), md)
			}
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)

	if err != nil {
		log.Fatalf("error connecting to server: %v", err)
	}

	return Client{
		AuthServiceClient:  monify.NewAuthServiceClient(conn),
		GroupServiceClient: monify.NewGroupServiceClient(conn),
		users:              &users,
		currentUser:        &currentUser,
	}
}

// CreateTestUser creates a new user and returns the userId
// and sets current user as the newly created user
func (c Client) CreateTestUser() string {
	email := uuid.New().String() + "@gmail.com"
	_, err := c.EmailRegister(context.TODO(), &monify.EmailRegisterRequest{Email: email, Password: "12345678"})
	if err != nil {
		log.Fatalf("error creating user: %v", err)
	}
	res, err := c.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email, Password: "12345678"})
	*c.currentUser = res.UserId
	(*c.users)[res.UserId] = res.AccessToken
	return res.UserId
}

// SetTestUser sets the current user to the user with the given userId
func (c Client) SetTestUser(userId string) {
	if _, ok := (*c.users)[userId]; !ok {
		log.Fatalf("user %s not found", userId)
	}
	*c.currentUser = userId
}
