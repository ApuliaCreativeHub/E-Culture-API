package utils

import (
	"E-Culture-API/utils"
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"io/ioutil"
	"os"
)

func MakeImgs(r io.Reader, path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return "", err
		}
	}

	all, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	randomString := utils.RandStringRunes(10)
	err = ioutil.WriteFile(path+"/"+randomString+"_n.png", all, 0655)
	if err != nil {
		return "", err
	}

	err = makeThumbnail(all, path, randomString)
	if err != nil {
		return "", err
	}
	return randomString, nil
}

func makeThumbnail(imgByte []byte, path string, name string) error {
	normalSizeImg, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}
	thumbnail := imaging.Thumbnail(normalSizeImg, 64, 64, imaging.Lanczos)

	err = imaging.Save(thumbnail, path+"/"+name+"_t.png")
	if err != nil {
		return err
	}
	return nil
}
