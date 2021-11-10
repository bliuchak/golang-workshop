package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()

	l.Info().Msg("start app aaa")

	r := chi.NewMux()
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})

	srv := &http.Server{
		Addr:    net.JoinHostPort("", "3000"),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		l.Fatal().Err(err).Msg("http server error")
	}
}
