package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cyber-missile/good-face-bad-face/internal/application"
	"github.com/cyber-missile/good-face-bad-face/internal/handler"
	"github.com/cyber-missile/good-face-bad-face/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func fileServer() (*http.Handler, error) {
	assets, err := web.Assets()
	if err != nil {
		return nil, err
	}

	fileServer := http.FileServer(http.FS(assets))

	return &fileServer, nil
}

func getRoutes(app *application.App) (*chi.Mux, error) {
	handlers := handler.New(app)
	router := chi.NewRouter()

	fileServer, err := fileServer()
	if err != nil {
		return nil, err
	}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logger(app.Logger.Sugar(), &Options{
		WithUserAgent: true,
		WithReferer:   true,
	}))
	router.Use(middleware.Recoverer)

	router.Handle("/static/*", http.StripPrefix("/static/", *fileServer))

	router.Get("/", handlers.Main)
	router.Get("/game/", handlers.Board)
	router.Get("/ws", handlers.Websocket)

	return router, nil
}

func Start(app *application.App, ctx context.Context) error {
	handler, err := getRoutes(app)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: handler,
	}

	app.Logger.Info("Starting http server")

	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}

		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
