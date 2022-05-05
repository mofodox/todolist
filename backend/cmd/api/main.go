package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apsdehal/go-logger"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *logger.Logger
}

func main() {
	var cfg config

	// Flag variables
	// e.g. We can change the port to 5000 with -port=5000 when running the server
	// For more information you can run `go run ./cmd/api -help`
	flag.IntVar(&cfg.port, "port", 4000, "API Server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production")
	flag.Parse()

	logger, err := logger.New(cfg.env, 1, os.Stdout)
	if err != nil {
		panic(err)
	}

	app := application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.InfoF("%s server is starting on port %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err.Error())
}
