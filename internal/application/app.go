package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cyber-missile/good-face-bad-face/internal/router"
)

type App struct {
	config string
}

func New(config string) *App {
	app := &App{
		config: config,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8000", //TODO: config
		Handler: router.GetRoutes(),
	}

	fmt.Println("Starting server")

	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}

		close(ch)
	}()

	var err error

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
