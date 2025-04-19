package main // Entry point of the application

import (
	"flag"
	"log"
	"os"

	"github.com/sklevenz/cf-api-server/internal/logger"
	"github.com/sklevenz/cf-api-server/internal/server"
)

func main() {
	var (
		port        = flag.Int("port", 8080, "HTTP server port")
		cfgDir      = flag.String("cfgDir", "./cfg", "Path to configuration directory")
		logFormat   = flag.String("logFormat", "text", "Log output format: text or json")
		logFilePath = flag.String("logFilePath", "./gen/cf-api-server.log", "Path to log file")
	)

	flag.Parse()

	jsonOut := false
	if *logFormat == "json" {
		jsonOut = true
	}

	file, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to create temp log file: %s %v", *logFilePath, err)
	}
	defer file.Close()

	logger.Log = logger.New(logger.LevelInfo, jsonOut, file)

	// Collect all flags into a map for logging
	fields := map[string]interface{}{}
	flag.VisitAll(func(f *flag.Flag) {
		fields[f.Name] = f.Value.String()
	})

	logger.Log.Info("Starting cf-api-server with settings: %v", fields)
	logger.Log.Info("Use -h or --help to display available flags.")

	server.StartServer(*port, *cfgDir)

}
