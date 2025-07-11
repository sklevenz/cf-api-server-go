package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"

	"github.com/sklevenz/cf-api-server/internal/gen"
	"github.com/sklevenz/cf-api-server/internal/logger"
	"github.com/sklevenz/cf-api-server/internal/testutil"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ gen.ServerInterface = (*Server)(nil)

type VersionInfo struct {
	SemanticVersion string `json:"semanticVersion"`
}
type Server struct {
	absCfgDir    string
	favicon      *[]byte
	rootDocument *gen.Root
	versionInfo  VersionInfo
	v3Document   *gen.V3
}

func NewServer(absCfgDir string, semver string) (*Server, error) {

	absCfgDir, err := filepath.Abs(absCfgDir)
	if err != nil {
		logger.Log.Error("Failed to set absCfgDir: %v", err)
		return nil, err
	}

	s := Server{
		absCfgDir: absCfgDir,
		versionInfo: VersionInfo{
			SemanticVersion: semver,
		},
	}
	s.loadFavicon()
	s.loadRootDocument()

	return &s, nil
}

func NewTestServer() (*Server, error) {
	absCfgDir, err := testutil.GetTestDataPath("cfg")
	if err != nil {
		logger.Log.Error("failed to get test data path: %v", err)
		return nil, err
	}
	s, err := NewServer(absCfgDir, "dev")
	return s, err
}

func (srv *Server) loadFavicon() {
	filePath := filepath.Join(srv.absCfgDir, "img", "favicon.ico")

	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Error("Failed to load favicon: %v", err)
		return
	}
	logger.Log.Info("Read favicon: %v", filePath)
	srv.favicon = &data
}

func (srv *Server) loadRootDocument() {
	filePath := filepath.Join(srv.absCfgDir, "template", "root.json.tmpl")
	logger.Log.Info("Read root document: %v", filePath)

	template, err := template.ParseFiles(filePath)
	if err != nil {
		logger.Log.Error("failed to parse root template: %v", err)
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, struct {
		BaseURL string
		Version string
	}{
		BaseURL: "http://localhost:8080",
		Version: srv.versionInfo.SemanticVersion,
	})
	if err != nil {
		logger.Log.Error("failed to render root template: %v", err)
	}
	logger.Log.Debug("After template processing: \n%s", &buf)

	var root gen.Root
	if err := json.Unmarshal(buf.Bytes(), &root); err != nil {
		logger.Log.Error("failed to unmarshal root JSON: %v", err)
	}

	logger.Log.Debug("Parsed root document: \n%+v", root)

	srv.rootDocument = &root
}

func (srv *Server) loadV3Document() {
	filePath := filepath.Join(srv.absCfgDir, "template", "v3.json.tmpl")
	logger.Log.Info("Read v3 document: %v", filePath)

	template, err := template.ParseFiles(filePath)
	if err != nil {
		logger.Log.Error("failed to parse v3 template: %v", err)
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, struct {
		BaseURL string
		Version string
	}{
		BaseURL: "http://localhost:8080",
		Version: srv.versionInfo.SemanticVersion,
	})
	if err != nil {
		logger.Log.Error("failed to render root template: %v", err)
	}
	logger.Log.Debug("After template processing: \n%s", &buf)

	var v3 gen.V3
	if err := json.Unmarshal(buf.Bytes(), &v3); err != nil {
		logger.Log.Error("failed to unmarshal root JSON: %v", err)
	}

	logger.Log.Debug("Parsed v3 document: \n%+v", v3)

	srv.v3Document = &v3
}
