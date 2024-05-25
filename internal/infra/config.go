package infra

type Config struct {
	PostgresURI string
}

func NewConfigFromSecrets(secrets map[string]string) Config {
	postgresUri := secrets["POSTGRES_URI"]
	if postgresUri == "" {
		panic("Missing POSTGRES_URI")
	}
	return Config{
		PostgresURI: postgresUri,
	}
}
