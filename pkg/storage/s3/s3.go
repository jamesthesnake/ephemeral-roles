package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ewohltman/ephemeral-roles/pkg/storage"
)

// Client is an S3 client that satisfies the storage.Client interface
type Client struct {
	bucket    string
	awsClient *s3.S3
}

// NewClient returns a new *Client associated with the given bucket name
func NewClient(bucketName string) (*Client, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	return &Client{
		bucket:    bucketName,
		awsClient: s3Client,
	}, nil
}

// List returns all the objects in the *Client bucket
func (client *Client) List() ([]string, error) {
	result, err := client.awsClient.ListObjects(
		&s3.ListObjectsInput{
			Bucket: &client.bucket,
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

// Store saves a Discord server custom configuration to the *Client bucket
func (client *Client) Store(config *storage.ServerConfig) {
	// TODO: Storing server config
}

// Retrieve returns a Discord server custom configuration from the *Client
// bucket
func (client *Client) Retrieve(server string) *storage.ServerConfig {
	// TODO: Retrieve server config
	return &storage.ServerConfig{}
}

// Delete removes a Discord server custom configuration from the *Client bucket
func (client *Client) Delete(server string) {
	// TODO: Delete server config
}
