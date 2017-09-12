package main

import (
	"github.com/therecipe/qt/core"
	"net/url"
)

type ImagePresenter struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *ImageModel `property:"imageModel"`

	_ func(path string)        `slot:"fileDropped"`
	_ func(row int) *ImageFile `slot:"getImage"`
}

func (p *ImagePresenter) init() {
	p.ConnectFileDropped(p.fileDropped)
	p.ConnectGetImage(p.getImage)
}

func (p *ImagePresenter) fileDropped(path string) {
	u, _ := url.Parse(path)
	var ir = new(ImageReader)
	images, _ := ir.ReadDir(u.Path)
	p.ImageModel().SetImageFile(images)
}

func (p *ImagePresenter) getImage(row int) *ImageFile {
	return p.ImageModel().ImageFile()[row]
}
