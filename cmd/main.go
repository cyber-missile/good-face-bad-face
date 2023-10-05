package main

import (
	"context"

	"github.com/cyber-missile/good-face-bad-face/internal/application"
)

func main() {
	app := application.New("")
	app.Start(context.TODO())
}
