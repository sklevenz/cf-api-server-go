package testutil

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetRepoRoot(t *testing.T) {
	root, err := GetRepoRoot()
	if err != nil {
		t.Fatalf("GetRepoRoot failed: %v", err)
	}

	// Check that path exists and is a directory
	info, err := os.Stat(root)
	if err != nil {
		t.Fatalf("Repo root path does not exist: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("Repo root path is not a directory: %s", root)
	}

	// Optional: check for presence of .git directory
	gitPath := filepath.Join(root, ".git")
	if _, err := os.Stat(gitPath); err != nil {
		t.Logf("Warning: .git directory not found at repo root: %s", gitPath)
	}
}

func TestGetTestDataPath(t *testing.T) {
	// This test assumes a file testdata/sample.txt exists
	testFile := "sample.txt"

	// Create it temporarily if it doesn't exist
	fullPath, err := GetTestDataPath(testFile)
	if err != nil {
		t.Fatalf("GetTestDataPath failed: %v", err)
	}

	// Create testdata/sample.txt if missing (non-intrusive test)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}
	content := "test content"
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	defer os.Remove(fullPath) // cleanup

	// Check the resolved path and file content
	data, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}
	if strings.TrimSpace(string(data)) != content {
		t.Errorf("Unexpected file content: %s", string(data))
	}
}
