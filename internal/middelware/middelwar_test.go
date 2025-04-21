package middleware_test

import (
	"os"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/logger"
)

func TestMain(m *testing.M) {
	logger.Log = logger.New(logger.LevelDebug, false, nil)
	os.Exit(m.Run())
}
