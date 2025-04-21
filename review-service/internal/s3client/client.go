package s3client

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	Client *s3.Client
	Bucket string
	Prefix string
}

func New(region, bucket, prefix string) (*S3Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	return &S3Service{
		Client: s3.NewFromConfig(cfg),
		Bucket: bucket,
		Prefix: prefix,
	}, nil
}

func (s *S3Service) ListJLFiles() ([]string, error) {
	var files []string
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(s.Prefix),
	}
	paginator := s3.NewListObjectsV2Paginator(s.Client, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, obj := range page.Contents {
			if strings.HasSuffix(*obj.Key, ".jl") {
				files = append(files, *obj.Key)
			}
		}
	}
	return files, nil
}

func (s *S3Service) DownloadFile(key string) ([]byte, error) {
	obj, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, obj.Body)
	return buf.Bytes(), err
}

