package mock

import (
	"context"
	"io"
)

type ObjectStore struct {
	PutFn func(ctx context.Context, key string, data io.Reader) error
}

func (os *ObjectStore) Put(ctx context.Context, key string, data io.Reader) error {
	return os.PutFn(ctx, key, data)
}
