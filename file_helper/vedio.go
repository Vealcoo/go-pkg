package filehelper

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"

	"github.com/pkg/errors"
)

func GetVideoThumbnail(data []byte, sizeOption ...int) (*bytes.Buffer, error) {
	size := 480
	if len(sizeOption) > 0 {
		size = sizeOption[0]
	}

	cmd := exec.Command("ffmpeg",
		"-i", "-",
		"-vframes", "1",
		"-vf", fmt.Sprintf("scale=w=%d:h=%d:force_original_aspect_ratio=decrease", size, size),
		//"-q:v", "3",
		"-f", "singlejpeg",
		"-",
	)

	cmd.Stdin = bytes.NewBuffer(data)

	buffer := new(bytes.Buffer)
	cmd.Stdout = buffer

	eBuf := new(bytes.Buffer)
	cmd.Stderr = eBuf

	if err := cmd.Run(); err != nil {
		if strings.Contains(eBuf.String(), "Cannot determine format of input stream") {
			return GetVideoFileThumbnail(data, size)
		}

		return nil, err
	} else if len(eBuf.Bytes()) > 0 {
		return GetVideoFileThumbnail(data, size)
	}

	return buffer, nil
}

func GetVideoFileThumbnail(data []byte, sizeOption ...int) (*bytes.Buffer, error) {
	size := 480
	if len(sizeOption) > 0 {
		size = sizeOption[0]
	}

	name, err := CreateTempFile(data)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("ffmpeg",
		"-i", name,
		"-vframes", "1",
		"-vf", fmt.Sprintf("scale=w=%d:h=%d:force_original_aspect_ratio=decrease", size, size),
		//"-q:v", "3",
		"-f", "singlejpeg",
		"-",
	)

	cmd.Stdin = bytes.NewBuffer(data)

	buffer := new(bytes.Buffer)
	cmd.Stdout = buffer

	eBuf := new(bytes.Buffer)
	cmd.Stderr = eBuf

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	errRemove := os.Remove(name)
	if errRemove != nil {

	}

	return buffer, nil
}

func CreateTempFile(data []byte) (string, error) {
	name := "/tmp/workdir/" + uuid.NewString()

	f, err := os.Create(name)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return name, nil
}
