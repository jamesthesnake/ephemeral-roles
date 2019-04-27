package s3

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewClient("ephemeral-roles-server-configs")
	if err != nil {
		t.Fatalf("Error creating new S3 client: %s", err)
	}
}

func TestService_List(t *testing.T) {
	s3Client, err := NewClient("ephemeral-roles-server-configs")
	if err != nil {
		t.Fatalf("Error creating new S3 client: %s", err)
	}

	_, err = s3Client.List()
	if err != nil {
		t.Fatalf("Error listing S3 bucket contents: %s", err)
	}
}
