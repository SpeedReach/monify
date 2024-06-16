package api

type ServerConfig struct {
	JwtSecret string
}

func NewConfigFromSecrets(secrets map[string]string) ServerConfig {
	return ServerConfig{
		JwtSecret: secrets["JWT_SECRET"],
	}
}
