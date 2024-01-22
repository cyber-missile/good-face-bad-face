package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cyber-missile/good-face-bad-face/internal/application"
	"github.com/cyber-missile/good-face-bad-face/internal/game"
	"github.com/cyber-missile/good-face-bad-face/internal/handler"
	"github.com/cyber-missile/good-face-bad-face/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func getRoutes(app *application.App, game *game.Game) *chi.Mux {
	handlers := handler.New(app, game)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logger(app.Logger.Sugar(), &Options{
		WithUserAgent: true,
		WithReferer:   true,
	}))
	router.Use(middleware.Recoverer)

	fileServer := http.FileServer(http.FS(web.StaticFilesDir))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	router.Get("/", handlers.Main)
	router.Get("/game/", handlers.Board)
	router.Get("/ws", handlers.NewRoom)
	router.Get("/ws/{roomUid}", handlers.EnterRoom)

	return router
}

func Start(app *application.App, game *game.Game, ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: getRoutes(app, game),
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
