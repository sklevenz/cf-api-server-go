package handler

import (
	"os"
	"testing"

	"github.com/sklevenz/cf-api-server/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.Log = logger.New(logger.LevelInfo, false, nil)
	os.Exit(m.Run())
}
