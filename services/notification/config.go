package notification

type Config struct {
	KafkaConn     string
	KafkaUser     string
	KafkaPassword string
}

func NewConfig(secrets map[string]string) Config {
	return Config{
		KafkaConn:     secrets["KAFKA_CONN"],
		KafkaUser:     secrets["KAFKA_USER"],
		KafkaPassword: secrets["KAFKA_PWD"],
	}
}
