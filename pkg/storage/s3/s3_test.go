package s3

import "testing"

func TestService_List(t *testing.T) {
	service := Service{}

	buckets, err := service.List()
	if err != nil {
		t.Fatalf("Error listing S3 buckets: %s", err)
	}

	t.Logf("Buckets: %v", buckets)
}
