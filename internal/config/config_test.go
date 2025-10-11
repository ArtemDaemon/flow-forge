package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempConfigFile(t *testing.T, content string) string {
	t.Helper()
	tmp := t.TempDir()
	file := filepath.Join(tmp, "config.yaml")
	if err := os.WriteFile(file, []byte(content), 0600); err != nil {
		t.Fatalf("failed to write temp config file: %v", err)
	}
	return file
}

func TestParseConfigFile_Success(t *testing.T) {
	yamlContent := `
input:
  file: "input.txt"
  type: "csv"
`
	file := writeTempConfigFile(t, yamlContent)
	cfg, err := ParseConfigFile(file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if cfg.Input.File != "input.txt" {
		t.Errorf("expected input.file to be 'input.txt', got '%s'", cfg.Input.File)
	}
	if cfg.Input.Type != "csv" {
		t.Errorf("expected input.type to be 'csv', got '%s'", cfg.Input.File)
	}
}

func TestParseConfigFile_MissingFileField(t *testing.T) {
	yamlContent := `
input:
  type: "csv"
`
	file := writeTempConfigFile(t, yamlContent)
	_, err := ParseConfigFile(file)
	if err == nil || err.Error() != "parsing config error: 'input.file' field is required" {
		t.Fatalf("expected error about missing 'input.file', got %v", err)
	}
}

func TestParseConfigFile_MissingTypeField(t *testing.T) {
	yamlContent := `
input:
  file: "input.txt"
`
	file := writeTempConfigFile(t, yamlContent)
	_, err := ParseConfigFile(file)
	if err == nil || err.Error() != "parsing config error: 'input.type' field is required" {
		t.Fatalf("expected error about missing 'input.type', got: %v", err)
	}
}

func TestParseConfigFile_WrongTypeFieldValue(t *testing.T) {
	yamlContent := `
input:
  file: "input.txt"
  type: "excel"
`
	file := writeTempConfigFile(t, yamlContent)
	_, err := ParseConfigFile(file)
	if err == nil || err.Error() != "parsing config error: 'input.type' field must have value 'csv'" {
		t.Fatalf("expected error about wrong 'input.type' value, got: %v", err)
	}
}

func TestParseConfigFile_FileDoesNotExist(t *testing.T) {
	_, err := ParseConfigFile("nonexistent.yaml")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestParseConfigFile_InvalidYAML(t *testing.T) {
	yamlContent := `
input:
  file: [not a string]
  type: 123
`
	file := writeTempConfigFile(t, yamlContent)
	_, err := ParseConfigFile(file)
	if err == nil {
		t.Fatal("expected error for invalid YAML, got nil")
	}
}
