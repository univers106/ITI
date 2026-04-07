package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"
)

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(string)
		expectPanic bool
		panicPrefix string
		expected    Config
	}{
		{
			name: "valid config",
			setup: func(path string) {
				content := `jwt_key: "my-secret-key"`
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatalf("failed to write test config: %v", err)
				}
			},
			expectPanic: false,
			expected:    Config{JwtKey: "my-secret-key"},
		},
		{
			name: "missing config creates example",
			setup: func(path string) {
				// no file
			},
			expectPanic: false,
			expected:    Config{JwtKey: "secret"},
		},
		{
			name: "invalid YAML",
			setup: func(path string) {
				content := `jwt_key: "unclosed quote`
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatalf("failed to write test config: %v", err)
				}
			},
			expectPanic: true,
			panicPrefix: "failed to unmarshal config:",
		},
		{
			name: "empty file",
			setup: func(path string) {
				if err := os.WriteFile(path, []byte(""), 0644); err != nil {
					t.Fatalf("failed to write test config: %v", err)
				}
			},
			expectPanic: false,
			expected:    Config{JwtKey: ""},
		},
		{
			name: "config missing field",
			setup: func(path string) {
				content := `other_field: "value"`
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatalf("failed to write test config: %v", err)
				}
			},
			expectPanic: false,
			expected:    Config{JwtKey: ""},
		},
		{
			name: "read error (directory instead of file)",
			setup: func(path string) {
				if err := os.Mkdir(path, 0755); err != nil {
					t.Fatalf("failed to create directory: %v", err)
				}
			},
			expectPanic: true,
			panicPrefix: "something wrong with config.yaml file:",
		},
		{
			name: "file does not exist and cannot create example",
			setup: func(path string) {
				dir := filepath.Dir(path)
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatalf("failed to create directory: %v", err)
				}
				if err := os.Chmod(dir, 0555); err != nil {
					t.Fatalf("failed to chmod directory: %v", err)
				}
			},
			expectPanic: true,
			panicPrefix: "failed to create example config:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			configPath := filepath.Join(tmpDir, "config.yaml")
			tt.setup(configPath)

			defer func() {
				r := recover()
				if tt.expectPanic {
					if r == nil {
						t.Error("expected panic but didn't panic")
						return
					}
					errMsg, ok := r.(string)
					if !ok {
						t.Errorf("panic value is not string: %v", r)
						return
					}
					if !strings.HasPrefix(errMsg, tt.panicPrefix) {
						t.Errorf("panic message %q does not have prefix %q", errMsg, tt.panicPrefix)
					}
				} else {
					if r != nil {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			config := ReadConfig(configPath)
			if !tt.expectPanic {
				if config.JwtKey != tt.expected.JwtKey {
					t.Errorf("expected JwtKey = %q, got %q", tt.expected.JwtKey, config.JwtKey)
				}
			}
		})
	}
}

func TestReadConfig_ExampleCreationFailsPanics(t *testing.T) {
	tmpDir := t.TempDir()
	readOnlyDir := filepath.Join(tmpDir, "readonly")
	if err := os.Mkdir(readOnlyDir, 0444); err != nil {
		t.Fatalf("failed to create read-only directory: %v", err)
	}
	configPath := filepath.Join(readOnlyDir, "config.yaml")

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic but didn't panic")
			return
		}
		errMsg, ok := r.(string)
		if !ok {
			t.Errorf("panic value is not string: %v", r)
			return
		}
		if !strings.HasPrefix(errMsg, "something wrong with config.yaml file:") {
			t.Errorf("unexpected panic message: %v", errMsg)
		}
	}()

	_ = ReadConfig(configPath)
}

func TestCreateExample(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "config.yaml")

		data, err := createExample(configPath)
		if err != nil {
			t.Errorf("createExample failed: %v", err)
		}
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Errorf("failed to read created file: %v", err)
		}
		if string(content) != string(data) {
			t.Errorf("file content doesn't match returned data. Got %q, expected %q", string(content), string(data))
		}

		expected := getExampleConfig(nil)
		if string(data) != string(expected) {
			t.Errorf("returned data doesn't match example config. Got %q, expected %q", string(data), string(expected))
		}
	})

	t.Run("directory not writable", func(t *testing.T) {
		tmpDir := t.TempDir()
		readOnlyDir := filepath.Join(tmpDir, "readonly")
		if err := os.Mkdir(readOnlyDir, 0444); err != nil {
			t.Fatalf("failed to create read-only directory: %v", err)
		}
		configPath := filepath.Join(readOnlyDir, "config.yaml")

		_, err := createExample(configPath)
		if err == nil {
			t.Error("expected error when directory is not writable")
		}
	})

	t.Run("file exists but not writable", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "config.yaml")
		// Create a read-only file
		if err := os.WriteFile(configPath, []byte("test"), 0444); err != nil {
			t.Fatalf("failed to create read-only file: %v", err)
		}

		_, err := createExample(configPath)
		if err == nil {
			t.Error("expected error when file is not writable")
		}
	})
}

func TestGetExampleConfig(t *testing.T) {
	data := getExampleConfig(nil)
	if len(data) == 0 {
		t.Error("getExampleConfig returned empty data")
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		t.Errorf("failed to unmarshal example config: %v", err)
	}
	if cfg.JwtKey != "secret" {
		t.Errorf("expected JwtKey = 'secret', got %q", cfg.JwtKey)
	}

	mock := &Config{JwtKey: "custom"}
	data = getExampleConfig(mock)
	var cfg2 Config
	if err := yaml.Unmarshal(data, &cfg2); err != nil {
		t.Errorf("failed to unmarshal mocked example config: %v", err)
	}
	if cfg2.JwtKey != "custom" {
		t.Errorf("expected JwtKey = 'custom', got %q", cfg2.JwtKey)
	}
}

func TestGetExampleConfig_PanicOnMarshalError(t *testing.T) {
	unmarshalable := make(chan int)

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic but didn't panic")
			return
		}
		errMsg, ok := r.(string)
		if !ok {
			t.Errorf("panic value is not string: %v", r)
			return
		}

		expectedMessages := []string{
			"cannot marshal example config",
			"cannot marshal type:",
		}
		found := false
		for _, expected := range expectedMessages {
			if strings.Contains(errMsg, expected) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("panic message should contain 'cannot marshal example config' or 'cannot marshal type:', got: %v", errMsg)
		}
	}()

	_ = getExampleConfig(unmarshalable)
}
