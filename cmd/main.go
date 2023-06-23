package main

import (
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	app, err := initServer(logger)
	if err != nil {
		logger.Fatal("unable to init server", zap.Error(err))
	}

	err = http.ListenAndServe(":8080", app)
	if err != nil {
		logger.Fatal("http server error", zap.Error(err))
	}
}
