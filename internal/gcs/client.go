package gcs

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
)

type Client struct {
	innerClient *storage.Client
	bucketName  string
}

func NewClient(ctx context.Context, bucketName string) (*Client, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		innerClient: client,
		bucketName:  bucketName,
	}, nil
}

func (c *Client) Put(ctx context.Context, key string, data io.Reader) error {
	bucket := c.innerClient.Bucket(c.bucketName)
	object := bucket.Object(key)

	w := object.NewWriter(ctx)
	_, err := io.Copy(w, data)
	if err != nil {
		return err
	}

	err = w.Close()
	return err
}
