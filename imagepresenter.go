package main

import (
	"github.com/therecipe/qt/core"
	"net/url"
)

type ImagePresenter struct {
	core.QObject

	pager *Pager

	_ func() `constructor:"init"`

	_ *ImageModel `property:"imageModel"`

	_ func(path string) `slot:"fileDropped"`
	_ func() int        `slot:"getPage"`
	_ func(page int)    `slot:"setPage"`
	_ func() bool       `slot:"hasNext"`
	_ func() bool       `slot:"hasPrev"`
	_ func() []string   `slot:"getImages"`
	_ func() []string   `slot:"nextImages"`
	_ func() []string   `slot:"prevImages"`
}

func (p *ImagePresenter) init() {
	p.ConnectFileDropped(p.fileDropped)
	p.ConnectGetPage(p.getPage)
	p.ConnectSetPage(p.setPage)
	p.ConnectHasNext(p.hasNext)
	p.ConnectHasPrev(p.hasPrev)
	p.ConnectGetImages(p.getImages)
	p.ConnectNextImages(p.nextImages)
	p.ConnectPrevImages(p.prevImages)

	p.pager = NewPager(0, 0)
}

func (p *ImagePresenter) fileDropped(path string) {
	u, _ := url.Parse(path)
	var ir = new(ImageReader)
	images, _ := ir.ReadDir(u.Path)
	p.ImageModel().SetImageFile(images)

	p.pager.CurrentPage = 1
	p.pager.LastPage = len(images)
}

func (p *ImagePresenter) getPage() int {
	return p.pager.CurrentPage
}

func (p *ImagePresenter) setPage(page int) {
	p.pager.SetCurrent(page)
}

func (p *ImagePresenter) hasNext() bool {
	if p.pager.LastPage > p.pager.CurrentPage {
		return true
	}
	return false
}

func (p *ImagePresenter) hasPrev() bool {
	if p.pager.CurrentPage > 1 {
		return true
	}
	return false
}

func (p *ImagePresenter) getImages() []string {
	right := p.getImage(p.pager.CurrentPage - 1)
	left := p.getImage(p.pager.CurrentPage)
	return []string{right.Original(), left.Original()}
}

func (p *ImagePresenter) nextImages() []string {
	p.pager.Next()

	return p.getImages()
}

func (p *ImagePresenter) prevImages() []string {
	p.pager.Prev()

	return p.getImages()
}

func (p *ImagePresenter) getImage(row int) *ImageFile {
	if p.pager.LastPage > row {
		return p.ImageModel().ImageFile()[row]
	}
	img := new(ImageFile)
	img.SetOriginal("")
	img.SetThumbnail("")
	return img
}
