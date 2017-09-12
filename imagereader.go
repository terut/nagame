package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ImageReader struct{}

func (r *ImageReader) ReadDir(dirName string) ([]*ImageFile, error) {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Errorf("Directory cannot read %s\n", err)
		return nil, err
	}

	var images []*ImageFile
	for _, fileInfo := range fileInfos {
		name := (fileInfo).Name()
		if strings.HasSuffix(name, ".jpg") {
			var path = "file://" + dirName + "/" + (fileInfo).Name()
			var image = NewImageFile(nil)
			image.SetOriginal(path)
			images = append(images, image)
		}
	}
	return images, nil
}

func IsDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	return fInfo.IsDir(), nil
}
