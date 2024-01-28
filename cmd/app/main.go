package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/happsie/go-webserver-template/internal/architecture"
	"github.com/happsie/go-webserver-template/internal/domain/user"
)

func main() {
	configPath := flag.String("config", "config.yml", "Specify config file")
	flag.Parse()
	jsonLogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	log := slog.New(jsonLogHandler)
	conf, err := architecture.LoadConfig(*configPath)
	if err != nil {
		log.Error("could not load config", "error", err)
		os.Exit(1)
	}
	db, err := architecture.InitDB(log, conf)
	if err != nil {
		log.Error("could not connect to database", "error", err)
		os.Exit(1)
	}
	c := &architecture.Container{
		DB:     db,
		Config: conf,
		L:      log,
	}
	r := architecture.Router{
		Port: 8080,
		RouteGroups: []architecture.Routes{
			user.InitAPI(c),
		},
	}
	if err := r.Start(); err != nil {
		log.Error("could not start http server", "error", err)
		os.Exit(1)
	}
}
