package gcs

import (
	"context"
	"privaTutle/model"

	"cloud.google.com/go/storage"
)

func Write(ctx context.Context, bucket *storage.BucketHandle, key string, val []byte) error {
	obj := bucket.Object(key)
	w := obj.NewWriter(ctx)

	if _, err := w.Write(val); err != nil {
		return model.ErrInternal
	}

	if err := w.Close(); err != nil {
		return model.ErrInternal
	}

	return nil
}

func Remove(ctx context.Context, bucket *storage.BucketHandle, key string) error {
	err := bucket.Object(key).Delete(ctx)
	if err != nil {
		err = model.ErrInternal
	}

	return err
}

func Copy(ctx context.Context, src *storage.ObjectHandle, dst *storage.ObjectHandle) error {
	_, err := dst.CopierFrom(src).Run(ctx)
	if err != nil {
		err = model.ErrInternal
	}

	return err
}
