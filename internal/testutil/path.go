package testutil

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetRepoRoot returns the absolute path to the Git repository root.
func GetRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// GetTestdataPath returns the full path to a file under testdata/
func GetTestDataPath(subpaths ...string) (string, error) {
	root, err := GetRepoRoot()
	if err != nil {
		return "", err
	}
	parts := append([]string{root, "testdata"}, subpaths...)
	return filepath.Join(parts...), nil
}
