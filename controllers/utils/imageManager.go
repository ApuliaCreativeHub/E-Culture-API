package utils

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"io/ioutil"
	"os"
)

func MakeImgs(r io.Reader, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}

	all, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path+"/normal_size.png", all, 0655)
	if err != nil {
		return err
	}

	err = makeThumbnail(all, path)
	if err != nil {
		return err
	}
	return nil
}

func makeThumbnail(imgByte []byte, path string) error {
	normalSizeImg, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}
	thumbnail := imaging.Thumbnail(normalSizeImg, 64, 64, imaging.Lanczos)

	err = imaging.Save(thumbnail, path+"/thumbnail.png")
	if err != nil {
		return err
	}
	return nil
}
