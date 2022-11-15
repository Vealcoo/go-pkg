package gcs

import (
	"context"

	"cloud.google.com/go/storage"
)

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
