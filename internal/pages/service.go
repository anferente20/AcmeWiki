package pages

import (
	"github.com/anferente20/AcmeWiki/internal/domain"
)

type PageService interface {
	/*  This method allows to save a new page*/
	Save(page domain.Page) error

	/* This method allos to load a single page */
	LoadPage(title string) (*domain.Page, error)
}

type PagServiceImpl struct {
	repo PageRepo
}

/*
Creates a PageService implementation
@param repository: A PageService Implementation
@return PageServiceImpl: the new page service
*/
func NewPageService(repository PageRepo) PageService {
	return &PagServiceImpl{repo: repository}
}

/*
Save a new page into a file. Calls his repo function
@param page: Page to be saved
@return error: if there is an error while writing the file
*/
func (psi *PagServiceImpl) Save(page domain.Page) error {
	return psi.repo.save(page)
}

/*
Get the page info by his title. Calls his repo function
@param title: Title of the page to be searched
@return *domain.Page: Pointer to the page retrieved
@return error: if there is an error searching the file
*/
func (psi *PagServiceImpl) LoadPage(title string) (*domain.Page, error) {
	return psi.repo.loadPage(title)
}
