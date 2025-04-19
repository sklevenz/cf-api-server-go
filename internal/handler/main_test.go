package handler

import (
	"os"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/logger"
)

// TestMain is the entry point for testing in this package. It sets up the
// logger with the desired configuration before running the test suite.
// The function takes a testing.M instance, which manages the execution
// of tests, and exits the program with the status code returned by m.Run().
func TestMain(m *testing.M) {
	logger.Log = logger.New(logger.LevelInfo, false, nil)
	os.Exit(m.Run())
}
