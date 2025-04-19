package handler

import (
	"os"
	"path/filepath"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/internal/logger"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ generated.ServerInterface = (*Server)(nil)

type Server struct {
	cfgDir  string
	favicon *[]byte
}

func NewServer(cfgDir string) Server {
	return Server{
		cfgDir:  cfgDir,
		favicon: nil,
	}
}

func (s *Server) LoadFavicon() {
	filePath := filepath.Join(s.cfgDir, "img", "favicon.ico")

	filePath, err := filepath.Abs(filePath)
	if err != nil {
		logger.Log.Error("Failed to load favicon: %v", err)
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Error("Failed to load favicon: %v", err)
		return
	}
	s.favicon = &data
}
