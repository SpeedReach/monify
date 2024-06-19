package media

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"monify/lib/utils"
	"net/http"
	"net/http/httptest"
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

func TestUpload(t *testing.T) {
	SetupTestServer()
	fp := getTestFilePath()
	file, err := os.Open(fp)
	assert.NoError(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	// Add the file to the form
	part, err := writer.CreateFormFile("image", filepath.Base("test.png"))
	assert.NoError(t, err)
	_, err = io.Copy(part, file)
	assert.NoError(t, err)
	err = writer.Close()
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/image", &requestBody)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a response recorder
	response := httptest.NewRecorder()
	state.server.mux.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

	resBody := UploadImageResponse{}
	err = json.Unmarshal(response.Body.Bytes(), &resBody)
	assert.NoError(t, err)
	println(resBody.Url)
	println(resBody.ImageId)
}
