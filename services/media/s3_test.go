package media

import (
	"github.com/stretchr/testify/assert"
	"monify/lib/utils"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

type TestServerState struct {
	server  Server
	infra   Infra
	started bool
	mutex   sync.Mutex
}

var state TestServerState

func SetupTestServer() {
	state.mutex.Lock()
	if state.started {
		state.mutex.Unlock()
		return
	}

	secrets, err := utils.LoadSecrets(utils.LoadEnv())
	if err != nil {
		panic(err)
	}
	infra, err := Setup(NewConfig("dev", secrets))
	if err != nil {
		panic(err)
	}
	state.infra = infra
	state.server = NewServer(infra)
	state.started = true
	state.mutex.Unlock()
}

func getTestFilePath() string {
	abs, err := filepath.Abs("test.png")
	if err != nil {
		panic(err)
	}
	return abs
}

func TestS3Storage(t *testing.T) {
	SetupTestServer()
	fp := getTestFilePath()
	file, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	err = state.infra.objStorage.Delete("test.png")
	assert.NoError(t, err)
	err = state.infra.objStorage.Store("test.png", file)
	assert.NoError(t, err)

	response, err := http.Get(state.infra.objStorage.GetUrl("test.png"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	err = state.infra.objStorage.Delete("test.png")
	assert.NoError(t, err)

	response, err = http.Get(state.infra.objStorage.GetUrl("test.png"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, response.StatusCode)
}
