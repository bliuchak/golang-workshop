package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bliuchak/golang-workshop/internal/api"
	"github.com/bliuchak/golang-workshop/internal/platform/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()

	l.Info().Msg("start app")

	db, err := storage.NewRedis("storage", "6379")
	if err != nil {
		l.Fatal().Err(err).Msg("redis init error")
	}

	rest := api.NewAPI(db, l)

	r := chi.NewMux()
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})

	r.Get("/book/{id}", rest.GetBook())
	r.Post("/book", rest.CreateBook())

	srv := &http.Server{
		Addr:    net.JoinHostPort("", "3000"),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		l.Fatal().Err(err).Msg("http server error")
	}
}
