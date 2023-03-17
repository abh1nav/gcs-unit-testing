package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/abh1nav/gcs-unit-testing/internal/gcs"
	"github.com/abh1nav/gcs-unit-testing/internal/giraffe"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	objectStore giraffe.ObjectStore
}

func main() {
	ctx := context.Background()
	app, err := buildApp(ctx)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/plain")

		data := strings.NewReader("hello world this is a text file")
		err := app.objectStore.Put(ctx, "upload-test/file.txt", data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	http.ListenAndServe("127.0.0.1:3000", r)
}

func buildApp(ctx context.Context) (*App, error) {
	gcsObjectStore, err := gcs.NewClient(ctx, "sandbox-abhi")
	if err != nil {
		return nil, err
	}

	return &App{
		objectStore: gcsObjectStore,
	}, nil
}
