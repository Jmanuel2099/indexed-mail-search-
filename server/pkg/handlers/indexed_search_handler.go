package handlers

import (
	"fmt"
	"indexed-mail-search/server/pkg/domain"
	"indexed-mail-search/server/pkg/handlers/contracts"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type IndexedSearchHAandler struct {
	indexedSearchService contracts.IIndexedSearch
}

func NewIndexedSearchHAandler(iss contracts.IIndexedSearch) *IndexedSearchHAandler {
	return &IndexedSearchHAandler{
		indexedSearchService: iss,
	}
}

type SearchTermInEmailsResponse struct {
	Emails []domain.Email `json:"emails"`
}

func (ish *IndexedSearchHAandler) SearchTermInEmails(w http.ResponseWriter, r *http.Request) {
	term := chi.URLParam(r, "term")

	emails, err := ish.indexedSearchService.SearchInIndexedEmails(term)
	if err != nil {
		// NewErrResponse(w, r, http.StatusInternalServerError, err)
		fmt.Print("err: " + err.Error())
		return
	}
	response := &SearchTermInEmailsResponse{
		Emails: emails,
	}
	fmt.Print(response.Emails)
	render.JSON(w, r, response)
}
