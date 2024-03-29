package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func gracefulServerShutdown(server *http.Server) {
	signalRx := make(chan os.Signal, 1)
	signal.Notify(signalRx, syscall.SIGINT, syscall.SIGTERM)

	signal := <-signalRx

	log.Info().Msgf("Signal received: %v", signal)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Error stopping server")
	}
}

func Server(port uint16) *http.Server {
	server := http.Server{
		Addr:              ":" + strconv.Itoa(int(port)),
		Handler:           nil,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	return &server
}

func RunServer(ctx context.Context, port uint16) error {
	server := Server(port)

	RegisterHandlers(server)

	go gracefulServerShutdown(server)

	log.Info().Msgf("Listening on port %d", port)
	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		log.Info().Msg("Server closed")
	} else if err != nil {
		log.Error().Err(err).Msg("Error starting server")
	}

	return nil
}
