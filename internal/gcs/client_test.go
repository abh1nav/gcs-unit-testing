package gcs

import (
	"context"
	"strings"
	"testing"

	"github.com/fsouza/fake-gcs-server/fakestorage"
)

func TestPut(t *testing.T) {
	testBucketName := "test-bucket"
	testData := "this is some test data"

	server := fakestorage.NewServer([]fakestorage.Object{
		{
			ObjectAttrs: fakestorage.ObjectAttrs{
				BucketName: testBucketName,
				Name:       "test-folder/test-file.txt",
			},
			Content: []byte("contents of test file"),
		},
	})
	defer server.Stop()
	innerClient := server.Client()

	gcsClient, err := NewClient(context.TODO(), testBucketName)
	if err != nil {
		t.Fail()
	}
	gcsClient.innerClient = innerClient

	err = gcsClient.Put(context.TODO(), "test-key/unit-test.txt", strings.NewReader(testData))
	if err != nil {
		t.Fail()
	}
}
