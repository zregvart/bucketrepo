package main

import (
	"flag"
	"fmt"
	"os"
)

type flags struct {
	logLevel   string
	configPath string
}

func main() {
	flags := flags{}
	flag.StringVar(&flags.logLevel, "log-level", "warning", "Defines the logs level (debug, info, warning, error, fatal, panic)")
	flag.StringVar(&flags.configPath, "config-path", ".", "Define the absolute path to the config.yaml file")
	flag.Parse()

	err := InitLogger(flags.logLevel)
	if err != nil {
		fmt.Printf("Invalid log-level option: %s", err)
		os.Exit(2)
	}

	config := NewConfig(flags.configPath)

	storage := NewStorage(config.Storage)
	cache := NewFileSystemStorage(config.Cache)
	repository := NewRepository(config.Repository)
	controller := NewFileController(cache, storage, repository)

	InitHTTP(config.HTTP, controller)
}
