package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type idLoginBody struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
type idLoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func GetAccessToken(clientId string, clientSecret string) (string, error) {
	url := "https://app.infisical.com/api/v1/auth/universal-auth/login"
	encoded, err := json.Marshal(&idLoginBody{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(encoded))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body := idLoginResponse{}
	bodyBytes, err := io.ReadAll(res.Body)
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return "", err
	}
	return body.AccessToken, nil
}

type Secrets struct {
	Secrets []Secret `json:"secrets"`
}
type Secret struct {
	SecretKey   string `json:"secretKey"`
	SecretValue string `json:"secretValue"`
}

func GetSecrets(token string, environment string) (Secrets, error) {
	url := fmt.Sprintf("https://app.infisical.com/api/v3/secrets/raw?workspaceId=6500618f61d12cd3d808036a&environment=%s&workspaceSlug=backend-6w3j", environment)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	secrets := Secrets{}
	err := json.NewDecoder(res.Body).Decode(&secrets)
	return secrets, err
}

func LoadSecrets(env string) (map[string]string, error) {
	println("Loading secrets with env: ", env)
	id, ok := os.LookupEnv("CLIENT_ID")
	if !ok {
		panic("cannot load client id for infisical")
	}
	s, ok := os.LookupEnv("CLIENT_SECRET")
	if !ok {
		panic("cannot load client id for infisical")
	}
	token, err := GetAccessToken(id, s)
	if err != nil {
		return map[string]string{}, err
	}
	secrets, err := GetSecrets(token, env)
	if err != nil {
		return map[string]string{}, err
	}
	secretsMap := make(map[string]string, len(secrets.Secrets))
	for _, s := range secrets.Secrets {
		secretsMap[s.SecretKey] = s.SecretValue
	}
	return secretsMap, nil
}

func LoadEnv() string {
	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return "dev"
	}
	switch env {
	case "dev":
		return env
	case "stage":
		return env
	case "production":
		return env
	default:
		return "dev"
	}
}
