package main

import (
	"URLShortener/internal/config"
	"URLShortener/internal/lib/logger/sl"
	"URLShortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	conf := config.MustLoad()

	log := config.SetupSlog(conf.Env)
	log.Info("URLShortener service started ", slog.String("env:", conf.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(conf.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	//TODO: init router: chi, chi render

	//TODO: run server:
}
