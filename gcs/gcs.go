package gcs

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
)

func GcsConn(credentials string) *storage.Client {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentials)
	gcsClient, err := storage.NewClient(context.Background())
	if err != nil {
		panic(err)
	}

	return gcsClient
}

func Write(ctx context.Context, bucket *storage.BucketHandle, key string, val []byte) error {
	obj := bucket.Object(key)
	w := obj.NewWriter(ctx)

	if _, err := w.Write(val); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func Remove(ctx context.Context, bucket *storage.BucketHandle, key string) error {
	err := bucket.Object(key).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func Copy(ctx context.Context, src *storage.ObjectHandle, dst *storage.ObjectHandle) error {
	_, err := dst.CopierFrom(src).Run(ctx)
	if err != nil {
		return err
	}

	return nil
}
