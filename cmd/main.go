package main

import (
	"context"
	"fmt"

	"github.com/cyber-missile/good-face-bad-face/internal/application"
	"github.com/cyber-missile/good-face-bad-face/internal/config"
	"github.com/cyber-missile/good-face-bad-face/internal/router"
	"github.com/cyber-missile/good-face-bad-face/internal/templates"
	"go.uber.org/zap"
)

func main() {
	err := start()
	if err != nil {
		fmt.Print(err)
	}
}

func start() error {
	config := config.Config{Port: 9001}

	templateCache, err := templates.New()
	if err != nil {
		return err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	defer logger.Sync()

	app := application.App{
		Config:    config,
		Templates: templateCache,
		Logger:    *logger,
	}

	ctx := context.Background()
	return router.Start(&app, ctx)
}
