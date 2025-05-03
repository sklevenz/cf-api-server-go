package main // Entry point of the application

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sklevenz/cf-api-server/internal/logger"
	"github.com/sklevenz/cf-api-server/internal/server"
)

var SemanticVersion = "dev"

func main() {
	var (
		port        = flag.Int("port", 8080, "HTTP server port")
		cfgDir      = flag.String("cfgDir", "./cfg", "Path to configuration directory")
		logFormat   = flag.String("logFormat", "text", "Log output format: text or json")
		logLevelStr = flag.String("logLevel", "INFO", "Log level: DEBUG, INFO, WARNING, ERROR")
		logFilePath = flag.String("logFilePath", "./logs/cf-api-server.log", "Path to log file")
		verbose     = flag.Bool("v", false, "Shortcut for -logLevel=DEBUG")
		logLevel    logger.LogLevel
	)

	flag.Parse()
	if *verbose {
		logLevel = logger.LevelDebug
	} else {
		logLevel = parseLogLevel(*logLevelStr)
	}

	jsonOut := false
	if *logFormat == "json" {
		jsonOut = true
	}

	file, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to create temp log file: %s %v", *logFilePath, err)
	}
	defer file.Close()

	logger.Log = logger.New(logLevel, jsonOut, file)

	// Collect all flags into a map for logging
	fields := map[string]interface{}{}
	flag.VisitAll(func(f *flag.Flag) {
		fields[f.Name] = f.Value.String()
	})

	logger.Log.Info("Starting cf-api-server with settings: %v", fields)
	logger.Log.Info("Version: %v", SemanticVersion)
	logger.Log.Info("Use -h or --help to display available flags.")

	server.StartServer(*port, *cfgDir, SemanticVersion)
}

func parseLogLevel(level string) logger.LogLevel {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return logger.LevelDebug
	case "INFO":
		return logger.LevelInfo
	case "WARNING", "WARN":
		return logger.LevelWarn
	case "ERROR":
		return logger.LevelError
	default:
		fmt.Fprintf(os.Stderr, "Unknown log level: %s, defaulting to INFO\n", level)
		return logger.LevelInfo
	}
}
