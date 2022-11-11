package filehelper

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"github.com/h2non/filetype"
	"github.com/pkg/errors"
)

func DownscaleImageDefault(data []byte, sizeOption ...int) (*bytes.Buffer, *bytes.Buffer, error) {
	size := 1080
	if len(sizeOption) > 0 {
		size = sizeOption[0]
	}

	thumbSize := 480
	if len(sizeOption) > 1 {
		thumbSize = sizeOption[1]
	}

	if typ, err := filetype.Match(data); err != nil {
		return nil, nil, errors.WithStack(err)
	} else if typ.Extension == "gif" {
		_, thumb, err := DownscaleImage(data, size, thumbSize)
		if err != nil {
			return nil, nil, err
		}

		return bytes.NewBuffer(data), thumb, nil
	}

	return DownscaleImage(data, size, thumbSize)
}

func DownscaleImage(data []byte, size int, thumbSize int) (*bytes.Buffer, *bytes.Buffer, error) {

	buf := bytes.NewBuffer(data)
	img, err := imaging.Decode(buf)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	buf, err = downscaleImage(img, size)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	thumbBuf, err := downscaleImage(img, thumbSize)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return buf, thumbBuf, nil
}

func downscaleImage(img image.Image, size int) (*bytes.Buffer, error) {
	rw, rh := getResizeWH(img, size)
	transImg := imaging.Resize(img, rw, rh, imaging.Lanczos)

	buf := new(bytes.Buffer)
	err := imaging.Encode(buf, transImg, imaging.JPEG, imaging.JPEGQuality(85))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return buf, nil
}

func getResizeWH(img image.Image, size int) (rw int, rh int) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	if w >= h {
		rw = w
	} else {
		rh = h
	}

	if rw > size {
		rw = size
	} else if rh > size {
		rh = size
	}

	return
}
