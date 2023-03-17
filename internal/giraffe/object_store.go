package giraffe

import (
	"context"
	"io"
)

type ObjectStore interface {
	Put(ctx context.Context, key string, data io.Reader) error
}
