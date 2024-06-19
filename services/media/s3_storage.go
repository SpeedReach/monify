package media

import (
	"bytes"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"monify/lib/media"
)

type S3MediaStorage struct {
	svc    *s3.S3
	config Config
	media.Storage
}

func (m S3MediaStorage) Store(path string, imageData []byte) error {
	_, err := m.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(m.config.S3Bucket),
		Key:    aws.String(path),
		Body:   bytes.NewReader(imageData),
	})
	if err != nil {
		return errors.Join(errors.New("could not upload to s3"), err)
	}
	return nil
}

func (m S3MediaStorage) Delete(path string) error {
	_, err := m.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(m.config.S3Bucket),
		Key:    aws.String(path),
	})
	return err
}

func (m S3MediaStorage) GetHost() string {
	return m.config.S3Host
}
func (m S3MediaStorage) GetUrl(path string) string {
	return m.config.S3Host + "/" + path
}

func NewS3MediaStorage(config Config) S3MediaStorage {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:                   credentials.NewStaticCredentials(config.S3KeyId, config.S3KeyValue, ""),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region:                        aws.String(config.S3Region),
	}))
	svc := s3.New(sess)
	return S3MediaStorage{
		svc:    svc,
		config: config,
	}
}
