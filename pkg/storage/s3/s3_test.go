package s3

import "testing"

func TestService_List(t *testing.T) {
	service := Bucket{
		Name: "ephemeral-roles-server-configs",
	}

	buckets, err := service.List()
	if err != nil {
		t.Fatalf("Error listing S3 bucket contents: %s", err)
	}

	t.Logf("Contents: %v", buckets)
}
