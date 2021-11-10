package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/bliuchak/golang-workshop/internal/platform/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type API struct {
	storage *storage.Redis
	logger  zerolog.Logger
}

func NewAPI(s *storage.Redis, l zerolog.Logger) *API {
	return &API{
		storage: s,
		logger:  l,
	}
}

func (a *API) GetBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "id")

		book, err := a.storage.GetBook(bookID)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
		}

		data, err := json.Marshal(book)
		if err != nil {
			a.logger.Error().Err(err).Msg("unable to marshall data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(data)
	}
}

func (a *API) CreateBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			a.logger.Error().Err(err).Msg("unable to read body")
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()

		var book storage.Book
		err = json.Unmarshal(b, &book)
		if err != nil {
			a.logger.Error().Err(err).Msg("unable to unmarshall data")
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = a.storage.CreateBook(book.ID, book.Title)
		if err != nil {
			a.logger.Error().Err(err).Msg("unable to create book")
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusNoContent)
	}
}
