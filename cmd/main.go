package main

import (
	"context"
	"fmt"

	"github.com/cyber-missile/good-face-bad-face/internal/application"
	"github.com/cyber-missile/good-face-bad-face/internal/config"
	"github.com/cyber-missile/good-face-bad-face/internal/router"
	"go.uber.org/zap"
)

func main() {
	err := start()
	if err != nil {
		fmt.Print(err)
	}
}

// Start dose start start
// and the start start tge start which start the start on start
func start() error {
	config := config.Config{Port: 9001}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	defer logger.Sync()

	app := application.App{
		Config: config,
		Logger: *logger,
	}

	ctx := context.Background()
	return router.Start(&app, ctx)
}
