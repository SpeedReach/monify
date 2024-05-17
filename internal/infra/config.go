package infra

type Config struct {
	PostgresURI string
}

func NewConfigFromSecrets(secrets map[string]string) Config {
	return Config{
		PostgresURI: secrets["POSTGRES_URI"],
	}
}
