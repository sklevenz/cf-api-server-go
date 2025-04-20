package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/internal/logger"
	"github.com/sklevenz/cf-api-server/internal/testutil"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ generated.ServerInterface = (*Server)(nil)

type VersionInfo struct {
	SemanticVersion string `json:"semanticVersion"`
}
type Server struct {
	cfgDir       string
	favicon      *[]byte
	rootDocument *generated.N200Root
	versionInfo  VersionInfo
}

func NewServer(cfgDir string, semver string) *Server {
	s := Server{
		cfgDir: cfgDir,
		versionInfo: VersionInfo{
			SemanticVersion: semver,
		},
	}
	s.loadFavicon()
	s.loadRootDocument()

	return &s
}

func NewTestServer() (*Server, error) {
	cfgDir, err := testutil.GetTestDataPath("cfg")
	if err != nil {
		logger.Log.Error("failed to get test data path: %v", err)
		return nil, err
	}
	s := NewServer(cfgDir, "dev")
	return s, nil
}

func (s *Server) loadFavicon() {
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

func (srv *Server) loadRootDocument() {
	templateFileName := filepath.Join(srv.cfgDir, "template", "root.json.tmpl")

	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		logger.Log.Error("failed to parse root template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		BaseURL string
		Version string
	}{
		BaseURL: "http://localhost:8080",
		Version: srv.versionInfo.SemanticVersion,
	})

	if err != nil {
		logger.Log.Error("failed to render root template: %v", err)
	}

	var root generated.N200Root
	if err := json.Unmarshal(buf.Bytes(), &root); err != nil {
		logger.Log.Error("failed to unmarshal root JSON: %v", err)
	}

	srv.rootDocument = &root
}
