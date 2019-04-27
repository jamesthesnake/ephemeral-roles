package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Service struct{}

func (service *Service) List() ([]string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	buckets := make([]string, 0)

	for _, bucket := range result.Buckets {
		buckets = append(buckets, *bucket.Name)
	}

	return buckets, nil
}
