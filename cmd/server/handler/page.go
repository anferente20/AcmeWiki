package handler

import (
	"github.com/anferente20/AcmeWiki/internal/domain"
	"github.com/anferente20/AcmeWiki/internal/pages"
)

type PageHandler struct {
	service pages.PageService
}

func NewPageHandler(ps pages.PageService) *PageHandler {
	return &PageHandler{
		service: ps,
	}
}

/*
Save a new page into a file. Calls his repo function
@param page: Page to be saved
@return error: if there is an error while writing the file
*/
func (ph *PageHandler) save(page domain.Page) error {
	return ph.service.Save(page)
}

/*
Get the page info by his title. Calls his repo function
@param title: Title of the page to be searched
@return *domain.Page: Pointer to the page retrieved
@return error: if there is an error searching the file
*/
func (ph *PageHandler) loadPage(title string) (*domain.Page, error) {
	return ph.service.LoadPage(title)
}
