package storeservice

import (
	"bytes"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
	"github.com/aws/aws-sdk-go/service/s3"
)

type (
	StoreService interface {
		GenerateUrl(keyName string) string
		PutObject(file multipart.File, fileHeader *multipart.FileHeader, id string) (string, error)
	}

	storeService struct {
		sess   *session.Session
		config *config.ConfigurationService
	}
)

func NewStoreService(sess *session.Session, config *config.ConfigurationService) StoreService {
	return &storeService{
		sess:   sess,
		config: config,
	}
}

func (s *storeService) PutObject(file multipart.File, fileHeader *multipart.FileHeader, id string) (string, error) {
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	tempFileName := "pictures/" + strings.Split(fileHeader.Filename, ".")[0] + id + filepath.Ext(fileHeader.Filename)
	_, err := s3.New(s.sess).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s.config.S3Config.AWSBucket),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

func (s *storeService) GenerateUrl(keyName string) string {
	req, _ := s3.New(s.sess).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.config.S3Config.AWSBucket),
		Key:    aws.String(keyName),
	})
	rest.Build(req)
	urlStr := req.HTTPRequest.URL.String()

	return urlStr
}
