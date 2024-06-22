package infra

type Config struct {
	PostgresURI    string
	KafkaUsername  string
	KafkaPassword  string
	KafkaConn      string
	FileServerHost string
}

func NewConfigFromSecrets(secrets map[string]string) Config {
	postgresUri := secrets["POSTGRES_URI"]
	kafkaU := secrets["KAFKA_USER"]
	kafkaP := secrets["KAFKA_PWD"]
	kafkaConn := secrets["KAFKA_CONN"]
	fileServerHost := secrets["FILE_SERVER_HOST"]

	if kafkaU == "" {
		panic("Missing KAFKA_USER")
	}
	if kafkaP == "" {
		panic("Missing KAFKA_PWD")
	}
	if kafkaConn == "" {
		panic("Missing KAFKA_CONN")
	}
	if postgresUri == "" {
		panic("Missing POSTGRES_URI")
	}
	if fileServerHost == "" {
		panic("Missing FILE_SERVER_HOST")
	}
	return Config{
		PostgresURI:    postgresUri,
		KafkaUsername:  kafkaU,
		KafkaPassword:  kafkaP,
		KafkaConn:      kafkaConn,
		FileServerHost: fileServerHost,
	}
}
