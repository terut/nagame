package main

func NewPager(currentPage int, lastPage int) *Pager {
	return &Pager{currentPage, lastPage}
}

type Pager struct {
	CurrentPage int
	LastPage    int
}

func (p *Pager) Next() {
	if p.LastPage < (p.CurrentPage + 2) {
		p.CurrentPage = p.LastPage
	} else {
		p.CurrentPage = p.CurrentPage + 2
	}
}

func (p *Pager) Prev() {
	if p.CurrentPage > 2 {
		p.CurrentPage = p.CurrentPage - 2
	} else {
		p.CurrentPage = 1
	}
}

func (p *Pager) SetCurrent(page int) {
	if page > 0 && page <= p.LastPage {
		p.CurrentPage = page
	}
}
