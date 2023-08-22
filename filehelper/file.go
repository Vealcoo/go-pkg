package filehelper

import (
	"io/ioutil"
	"mime/multipart"
)

func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func IsImage(in string) bool {
	switch in {
	case "image/bmp":
		fallthrough
	case "image/gif":
		fallthrough
	case "image/webp":
		fallthrough
	case "image/png":
		fallthrough
	case "image/jpeg":
		return true
	default:
		return false
	}
}

func IsVideo(in string) bool {
	switch in {
	case "video/mp4":
		return true
	default:
		return false
	}
}
