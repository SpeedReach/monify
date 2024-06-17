package media

type Config struct {
	PostgresURI string
	S3Host      string
	S3Bucket    string
	S3KeyId     string
	S3KeyValue  string
	S3Region    string
	JwtSecret   string
}

func NewConfig(secrets map[string]string) Config {
	return Config{
		JwtSecret:   secrets["JWT_SECRET"],
		PostgresURI: secrets["POSTGRES_URI"],
		S3Host:      secrets["S3_HOST"],
		S3Bucket:    secrets["S3_BUCKET"],
		S3KeyId:     secrets["S3_KEY_ID"],
		S3KeyValue:  secrets["S3_KEY_VALUE"],
		S3Region:    secrets["S3_REGION"],
	}
}
