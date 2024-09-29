package objectstore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"log"
)

var Session *session.Session

func initS3Session() error {
	var err error
	configService := config.GetConfigService()
	Session, err = session.NewSession(&aws.Config{
		Region: aws.String(configService.S3Config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			configService.S3Config.AWSAccessKeyID,
			configService.S3Config.AWSSecretAccessKey,
			""), // token can be left blank for now
	})
	if err != nil {
		log.Printf("storeservice: error initializing s3 session: %v", err)
		return err
	}

	return nil
}

func GetS3SessionInstance() (*session.Session, error) {
	if Session == nil {
		if err := initS3Session(); err != nil {
			return nil, err
		}
	}

	return Session, nil
}
