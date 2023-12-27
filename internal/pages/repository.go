package pages

import (
	"os"

	"github.com/anferente20/AcmeWiki/internal/domain"
)

type PageRepo interface {
	/*  This method allows to save a new page*/
	save(page domain.Page) error

	/* This method allos to load a single page */
	loadPage(title string) (*domain.Page, error)
}

type PageRepoImpl struct {
}

/*
Creates a PageRepo implementation
@return PageRepoImpl: the new page repo
*/
func NewPageRepository() PageRepo {
	return &PageRepoImpl{}
}

/*
Save a new page into a file
@param page: Page to be saved
@return error: if there is an error while writing the file
*/
func (pri *PageRepoImpl) save(page domain.Page) error {
	filepath := page.Title + ".txt"
	return os.WriteFile(filepath, page.Body, 0600)
}

/*
Get the page info by his title.
@param title: Title of the page to be searched
@return *domain.Page: Pointer to the page retrieved
@return error: if there is an error searching the file
*/
func (pri *PageRepoImpl) loadPage(title string) (*domain.Page, error) {
	filepath := title + ".txt"
	body, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return &domain.Page{Title: title, Body: body}, nil
}
