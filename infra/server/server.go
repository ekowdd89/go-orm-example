package server

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	Handler http.Handler
	Address string
}

func NewServer(svr ServerConfig) *http.Server {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf("listening on %s", svr.Address)
	return &http.Server{
		Handler: svr.Handler,
		Addr:    svr.Address,
	}
}
