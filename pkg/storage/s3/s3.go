package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Bucket struct {
	Name string
}

func (bucket *Bucket) List() ([]string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	result, err := s3Client.ListObjects(
		&s3.ListObjectsInput{
			Bucket: &bucket.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	items := make([]string, 0)

	for _, item := range result.Contents {
		items = append(items, *item.Key)
	}

	return items, nil
}
