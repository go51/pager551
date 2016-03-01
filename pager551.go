package pager551

import "strconv"

type pager struct {
	AllItems     int64
	AllPage      int64
	ItemsPerPage int64
	Page         int64
	Prev         *pageItem
	Next         *pageItem
	Pages        []*pageItem
	minPage      int64
	maxPage      int64
}

type pageItem struct {
	Disable bool
	Active  bool
	Label   string
	Page    int64
}

func New(allItems, itemsPerPage, page int64) *pager {
	pager := &pager{
		AllItems:     allItems,
		ItemsPerPage: itemsPerPage,
		Page:         page,
	}

	pager.load()

	return pager
}

func (p *pager) load() {
	if p.AllItems <= p.ItemsPerPage {
		p.noPageItems()
		return
	}

	p.AllPage = p.AllItems / p.ItemsPerPage
	if p.AllItems%p.ItemsPerPage != 0 {
		p.AllPage++
	}

	if p.AllPage > 5 {
		p.minPage = p.Page - 2
		p.maxPage = p.Page + 2
		switch {
		case p.minPage < 1:
			p.minPage = 1
			p.maxPage = p.minPage + 4
		case p.maxPage > p.AllPage:
			p.maxPage = p.AllPage
			p.minPage = p.maxPage - 4
		}
	} else {
		p.minPage = 1
		p.maxPage = p.AllPage
	}

	p.setPageItems()

}

func (p *pager) noPageItems() {
	prev := prevPageItem()
	next := nextPageItem()
	page := pagePageItem(1)

	prev.Disable = true
	next.Disable = true
	page.Active = true

	p.Prev = prev
	p.Next = next
	p.Pages = append(p.Pages, page)

}

func (p *pager) setPageItems() {
	p.setPrevItem()
	p.setNextItem()
	p.setPageItem()

}

func (p *pager) setPrevItem() {
	prev := prevPageItem()

	if p.Page <= 1 {
		prev.Disable = true
	} else {
		prev.Page = p.Page - 1
	}

	p.Prev = prev

}

func (p *pager) setNextItem() {
	next := nextPageItem()

	if p.maxPage <= p.Page {
		next.Disable = true
	} else {
		next.Page = p.Page + 1
	}

	p.Next = next

}

func (p *pager) setPageItem() {
	for i := p.minPage; i <= p.maxPage; i++ {
		page := pagePageItem(i)
		if i == p.Page {
			page.Active = true
		}
		p.Pages = append(p.Pages, page)
	}
}

func prevPageItem() *pageItem {
	return &pageItem{
		Disable: false,
		Active:  false,
		Label:   "&laquo;",
		Page:    0,
	}
}

func nextPageItem() *pageItem {
	return &pageItem{
		Disable: false,
		Active:  false,
		Label:   "&raquo;",
		Page:    0,
	}
}

func pagePageItem(page int64) *pageItem {
	return &pageItem{
		Disable: false,
		Active:  false,
		Label:   strconv.FormatInt(page, 10),
		Page:    page,
	}
}
