package main

import (
	"bytes"
	"encoding/json"
	"io"
	monify "monify/protobuf/gen/go"
	"net/http"
	"testing"
)

func TestWWWW(t *testing.T) {
	req1 := monify.EmailRegisterRequest{Email: "string", Password: "string"}
	body, _ := json.Marshal(&req1)
	res, _ := http.Post("http://api.monify.dev:8081/v1/user/register", "application/json", bytes.NewBuffer(body))
	parsedBody := monify.EmailRegisterResponse{}
	req, _ := io.ReadAll(res.Body)
	json.Unmarshal(req, &parsedBody)
	println(parsedBody.UserId)

	request := monify.EmailLoginRequest{Email: "string", Password: "string"}
	body, _ = json.Marshal(&request)
	post, err := http.Post("http://api.monify.dev:8081/v1/user/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return
	}
	req, _ = io.ReadAll(post.Body)
	println(post.StatusCode)
	println(string(req))
}
