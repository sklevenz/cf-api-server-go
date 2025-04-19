package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogger_Info_TextOutput(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LevelInfo, false, nil)
	logr.writer = &buf // Redirect output to buffer for testing

	logr.Info("Hello World %s", "paul")

	output := buf.String()

	if !strings.Contains(output, "[INFO] Hello World") {
		t.Errorf("Expected INFO log message, got: %s", output)
	}

	if !strings.Contains(output, "paul") {
		t.Errorf("Expected field user=paul in log, got: %s", output)
	}
}

func TestLogger_Info_JSONOutput(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LevelInfo, true, nil)
	logr.writer = &buf // Redirect output to buffer for testing

	logr.Info("Hello JSON %v", map[string]interface{}{
		"user": "paul",
	})

	output := buf.String()

	if !strings.Contains(output, `"level":"INFO"`) {
		t.Errorf("Expected JSON level INFO, got: %s", output)
	}

	if !strings.Contains(output, `"msg":"Hello JSON`) {
		t.Errorf("Expected JSON msg Hello JSON, got: %s", output)
	}

	if !strings.Contains(output, `user:paul`) {
		t.Errorf("Expected JSON field user=paul, got: %s", output)
	}
}
