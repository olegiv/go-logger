package logger

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestNew(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:      "info",
		LogDir:     tmpDir,
		MaxSizeMB:  10,
		MaxBackups: 5,
		Console:    false,
	}

	logger := New(cfg)

	if logger == nil {
		t.Fatal("Expected logger to be created")
	}
}

func TestNew_WithDefaults(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		LogDir: tmpDir,
	}

	logger := New(cfg)

	if logger == nil {
		t.Fatal("Expected logger to be created with defaults")
	}
}

func TestNew_WithConsole(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:      "debug",
		LogDir:     tmpDir,
		MaxSizeMB:  10,
		MaxBackups: 5,
		Console:    true,
	}

	logger := New(cfg)

	if logger == nil {
		t.Fatal("Expected logger to be created with console output")
	}
}

func TestNew_InvalidDirectory(t *testing.T) {
	// Use a path that likely can't be created (root-owned or invalid)
	cfg := Config{
		Level:      "info",
		LogDir:     "/this/path/should/not/exist/and/fail",
		MaxSizeMB:  10,
		MaxBackups: 5,
		Console:    false,
	}

	logger := New(cfg)

	// Should still create logger (fallback to stderr)
	if logger == nil {
		t.Fatal("Expected logger to be created even with invalid directory (fallback)")
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		expected zerolog.Level
	}{
		{"Debug", "debug", zerolog.DebugLevel},
		{"Info", "info", zerolog.InfoLevel},
		{"Warn", "warn", zerolog.WarnLevel},
		{"Warning", "warning", zerolog.WarnLevel},
		{"Error", "error", zerolog.ErrorLevel},
		{"Debug uppercase", "DEBUG", zerolog.DebugLevel},
		{"Info mixed case", "Info", zerolog.InfoLevel},
		{"Unknown", "unknown", zerolog.InfoLevel},
		{"Empty", "", zerolog.InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseLogLevel(tt.level)
			if result != tt.expected {
				t.Errorf("Expected level %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestClose(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	if logger == nil {
		t.Fatal("Expected logger to be created")
	}

	err := logger.Close()

	if err != nil {
		t.Errorf("Close should not return error: %v", err)
	}
}

func TestWithField(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	newLogger := logger.WithField("test_key", "test_value")

	if newLogger == nil {
		t.Fatal("Expected logger with field")
	}

	// Verify it returns a new logger instance
	if newLogger == logger {
		t.Error("WithField should return a new logger instance")
	}
}

func TestWithFields(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	fields := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
	}

	newLogger := logger.WithFields(fields)

	if newLogger == nil {
		t.Fatal("Expected logger with fields")
	}

	// Verify it returns a new logger instance
	if newLogger == logger {
		t.Error("WithFields should return a new logger instance")
	}
}

func TestWithFields_EmptyMap(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	fields := map[string]interface{}{}

	newLogger := logger.WithFields(fields)

	if newLogger == nil {
		t.Fatal("Expected logger even with empty fields")
	}
}

func TestWithError(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	err := errors.New("test error")
	newLogger := logger.WithError(err)

	if newLogger == nil {
		t.Fatal("Expected logger with error")
	}

	// Verify it returns a new logger instance
	if newLogger == logger {
		t.Error("WithError should return a new logger instance")
	}
}

func TestLogFileCreation(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:      "info",
		LogDir:     tmpDir,
		Filename:   "test.log",
		MaxSizeMB:  10,
		MaxBackups: 5,
		Console:    false,
	}

	logger := New(cfg)

	// Write a log message
	logger.Info().Msg("Test log message")

	// Check that log file was created
	logFile := filepath.Join(tmpDir, "test.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file should be created")
	}
}

func TestConfigStructure(t *testing.T) {
	cfg := Config{
		Level:      "debug",
		LogDir:     "/tmp/logs",
		Filename:   "custom.log",
		MaxSizeMB:  20,
		MaxBackups: 10,
		Console:    true,
	}

	if cfg.Level != "debug" {
		t.Error("Level not set correctly")
	}

	if cfg.LogDir != "/tmp/logs" {
		t.Error("LogDir not set correctly")
	}

	if cfg.Filename != "custom.log" {
		t.Error("Filename not set correctly")
	}

	if cfg.MaxSizeMB != 20 {
		t.Error("MaxSizeMB not set correctly")
	}

	if cfg.MaxBackups != 10 {
		t.Error("MaxBackups not set correctly")
	}

	if !cfg.Console {
		t.Error("Console not set correctly")
	}
}

func TestDefaultValues(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with empty config
	cfg := Config{
		LogDir: tmpDir,
	}

	logger := New(cfg)

	if logger == nil {
		t.Fatal("Expected logger to be created with defaults")
	}

	// Verify log file is created with defaults
	logger.Info().Msg("Test")
	logFile := filepath.Join(tmpDir, "go.log")

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file should be created with default settings")
	}
}

func TestAllLogLevels(t *testing.T) {
	tmpDir := t.TempDir()

	levels := []string{"debug", "info", "warn", "error"}

	for _, level := range levels {
		t.Run(level, func(t *testing.T) {
			cfg := Config{
				Level:  level,
				LogDir: filepath.Join(tmpDir, level),
			}

			logger := New(cfg)
			if logger == nil {
				t.Fatalf("Expected logger with level %s", level)
			}

			// Write a log message at each level
			logger.Debug().Msg("Debug message")
			logger.Info().Msg("Info message")
			logger.Warn().Msg("Warn message")
			logger.Error().Msg("Error message")
		})
	}
}

func TestWithMultipleFields(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)

	// Chain multiple WithField calls
	logger = logger.WithField("field1", "value1")
	logger = logger.WithField("field2", 42)
	logger = logger.WithField("field3", true)

	if logger == nil {
		t.Fatal("Expected logger after chaining WithField")
	}
}

func TestLoggerInheritance(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)

	// Verify Logger wraps zerolog.Logger
	// zerolog levels can be negative (e.g., Trace is -1), so just verify logger is functional
	if logger == nil {
		t.Fatal("Expected logger to be created")
	}

	// Test that the logger can log
	logger.Info().Msg("Test inheritance")
}

func TestConsoleOutput(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with console enabled
	cfg := Config{
		Level:   "debug",
		LogDir:  tmpDir,
		Console: true,
	}

	logger := New(cfg)
	logger.Info().Msg("Test console output")

	// Verify log file is still created
	logFile := filepath.Join(tmpDir, "go.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file should be created even with console enabled")
	}
}

func TestLogRotationSettings(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:      "info",
		LogDir:     tmpDir,
		MaxSizeMB:  50,
		MaxBackups: 10,
	}

	logger := New(cfg)

	// Write multiple log messages
	for i := 0; i < 100; i++ {
		logger.Info().Int("iteration", i).Msg("Test message")
	}

	// Verify log file exists
	logFile := filepath.Join(tmpDir, "go.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file should exist")
	}
}

func TestWithFieldsPreservesOriginal(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	original := New(cfg)
	modified := original.WithField("test", "value")

	// Original should be unchanged
	if original == modified {
		t.Error("WithField should create a new logger instance")
	}
}

func TestWithErrorNilError(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	newLogger := logger.WithError(nil)

	if newLogger == nil {
		t.Fatal("Expected logger even with nil error")
	}
}

func TestLogDirCreation(t *testing.T) {
	tmpDir := t.TempDir()
	nestedDir := filepath.Join(tmpDir, "nested", "log", "dir")

	cfg := Config{
		Level:  "info",
		LogDir: nestedDir,
	}

	logger := New(cfg)

	// Verify nested directory was created
	if _, err := os.Stat(nestedDir); os.IsNotExist(err) {
		t.Error("Nested log directory should be created")
	}

	// Verify logger works
	logger.Info().Msg("Test")
}

func TestEmptyLogDir(t *testing.T) {
	cfg := Config{
		Level:  "info",
		LogDir: "",
	}

	logger := New(cfg)

	if logger == nil {
		t.Fatal("Expected logger with empty LogDir (should use default)")
	}

	// Default should be "./logs"
	// Verify it was created or fallback to stderr
	logger.Info().Msg("Test")
}

func TestFilenameDefault(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:  "info",
		LogDir: tmpDir,
	}

	logger := New(cfg)
	logger.Info().Msg("Test default filename")

	// Verify default filename "go.log" is created
	logFile := filepath.Join(tmpDir, "go.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file should be created with default filename 'go.log'")
	}
}

func TestCustomFilename(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := Config{
		Level:    "info",
		LogDir:   tmpDir,
		Filename: "myapp.log",
	}

	logger := New(cfg)
	logger.Info().Msg("Test custom filename")

	// Verify custom filename is used
	logFile := filepath.Join(tmpDir, "myapp.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file should be created with custom filename 'myapp.log'")
	}

	// Verify default filename is NOT created
	defaultFile := filepath.Join(tmpDir, "go.log")
	if _, err := os.Stat(defaultFile); !os.IsNotExist(err) {
		t.Error("Default log file should not be created when custom filename is specified")
	}
}

func TestDirModeDefault(t *testing.T) {
	tmpDir := t.TempDir()
	logDir := filepath.Join(tmpDir, "test-perms")

	cfg := Config{
		LogDir: logDir,
	}

	logger := New(cfg)
	if logger == nil {
		t.Fatal("Expected logger to be created")
	}

	// Verify directory was created with default permissions (0750)
	info, err := os.Stat(logDir)
	if err != nil {
		t.Fatalf("Log directory should be created: %v", err)
	}

	expectedPerms := os.FileMode(0750)
	actualPerms := info.Mode().Perm()
	if actualPerms != expectedPerms {
		t.Errorf("Expected directory permissions %o, got %o", expectedPerms, actualPerms)
	}
}

func TestDirModeCustom(t *testing.T) {
	tmpDir := t.TempDir()
	logDir := filepath.Join(tmpDir, "test-custom-perms")

	cfg := Config{
		LogDir:  logDir,
		DirMode: 0700,
	}

	logger := New(cfg)
	if logger == nil {
		t.Fatal("Expected logger to be created")
	}

	// Verify directory was created with custom permissions (0700)
	info, err := os.Stat(logDir)
	if err != nil {
		t.Fatalf("Log directory should be created: %v", err)
	}

	expectedPerms := os.FileMode(0700)
	actualPerms := info.Mode().Perm()
	if actualPerms != expectedPerms {
		t.Errorf("Expected directory permissions %o, got %o", expectedPerms, actualPerms)
	}
}

func TestPathTraversalPrevention(t *testing.T) {
	tests := []struct {
		name           string
		logDir         string
		filename       string
		shouldFallback bool
		description    string
	}{
		{
			name:           "simple path traversal",
			logDir:         "../../etc",
			filename:       "test.log",
			shouldFallback: true,
			description:    "Should block simple .. traversal",
		},
		{
			name:           "nested path traversal",
			logDir:         "./logs/../../etc",
			filename:       "test.log",
			shouldFallback: true,
			description:    "Should block nested traversal",
		},
		{
			name:           "filename with path separator",
			logDir:         "./logs",
			filename:       "../etc/test.log",
			shouldFallback: true,
			description:    "Should block filename with path separator",
		},
		{
			name:           "filename with backslash",
			logDir:         "./logs",
			filename:       "..\\etc\\test.log",
			shouldFallback: true,
			description:    "Should block filename with backslash",
		},
		{
			name:           "valid relative path",
			logDir:         "./logs/app",
			filename:       "test.log",
			shouldFallback: false,
			description:    "Should allow valid relative path",
		},
		{
			name:           "valid nested directory",
			logDir:         "./logs/level1/level2",
			filename:       "test.log",
			shouldFallback: false,
			description:    "Should allow valid nested directory",
		},
		{
			name:           "filename with dots but not traversal",
			logDir:         "./logs",
			filename:       "app.v2.log",
			shouldFallback: false,
			description:    "Should allow filename with dots (not traversal)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			// For valid paths, use tmpDir as base
			logDir := tt.logDir
			if !tt.shouldFallback {
				logDir = filepath.Join(tmpDir, tt.logDir)
			}

			cfg := Config{
				LogDir:   logDir,
				Filename: tt.filename,
			}

			logger := New(cfg)
			if logger == nil {
				t.Fatal("Expected logger to be created (should fall back to stderr if needed)")
			}

			// For paths that should be blocked, verify directory wasn't created outside tmpDir
			if tt.shouldFallback {
				// Logger should have fallen back to stderr
				// We can't directly check if it's using stderr, but we can verify
				// that the dangerous directory wasn't created
				if strings.Contains(tt.logDir, "..") {
					// Directory with traversal shouldn't exist
					if _, err := os.Stat(tt.logDir); err == nil {
						t.Errorf("%s: dangerous directory was created: %s", tt.description, tt.logDir)
					}
				}
			} else {
				// For valid paths, verify the directory was created
				if _, err := os.Stat(logDir); os.IsNotExist(err) {
					t.Errorf("%s: valid directory was not created: %s", tt.description, logDir)
				}
			}
		})
	}
}

func TestPathCleaning(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name       string
		inputDir   string
		shouldWork bool
	}{
		{
			name:       "path with redundant slashes",
			inputDir:   "./logs//app///debug",
			shouldWork: true,
		},
		{
			name:       "path with dot segments",
			inputDir:   "./logs/./app/./debug",
			shouldWork: true,
		},
		{
			name:       "mixed separators (cleaned)",
			inputDir:   "./logs/app/debug",
			shouldWork: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logDir := filepath.Join(tmpDir, tt.inputDir)

			cfg := Config{
				LogDir: logDir,
			}

			logger := New(cfg)
			if logger == nil {
				t.Fatal("Expected logger to be created")
			}

			if tt.shouldWork {
				// Verify directory exists (path should be cleaned and created)
				if _, err := os.Stat(logDir); os.IsNotExist(err) {
					t.Errorf("Directory should exist after path cleaning: %s", logDir)
				}
			}
		})
	}
}

func TestNoGlobalLevelPollution(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a baseline zerolog logger at DEBUG level
	baselineLogger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)

	// Verify baseline logger is at DEBUG level
	if baselineLogger.GetLevel() != zerolog.DebugLevel {
		t.Fatalf("Baseline logger should be at DEBUG level, got %v", baselineLogger.GetLevel())
	}

	// Create go-logger instance at ERROR level
	cfg := Config{
		LogDir: tmpDir,
		Level:  "error",
	}
	errorLogger := New(cfg)

	// Verify go-logger has ERROR level
	if errorLogger.GetLevel() != zerolog.ErrorLevel {
		t.Errorf("go-logger should be at ERROR level, got %v", errorLogger.GetLevel())
	}

	// CRITICAL: Verify baseline logger still has DEBUG level (not polluted)
	if baselineLogger.GetLevel() != zerolog.DebugLevel {
		t.Errorf("Baseline logger level should remain DEBUG, got %v (global state pollution detected!)", baselineLogger.GetLevel())
	}

	// Verify global level was not changed
	globalLevel := zerolog.GlobalLevel()
	// Note: zerolog.GlobalLevel() defaults to TraceLevel if never set
	// We just verify that creating our logger didn't force it to ErrorLevel
	if globalLevel == zerolog.ErrorLevel {
		t.Errorf("Global level should not be ERROR (indicates SetGlobalLevel was called)")
	}
}

func TestMultipleLoggersWithDifferentLevels(t *testing.T) {
	tmpDir := t.TempDir()

	// Create logger at DEBUG level
	cfg1 := Config{
		LogDir: filepath.Join(tmpDir, "debug"),
		Level:  "debug",
	}
	debugLogger := New(cfg1)

	// Create logger at ERROR level
	cfg2 := Config{
		LogDir: filepath.Join(tmpDir, "error"),
		Level:  "error",
	}
	errorLogger := New(cfg2)

	// Verify both loggers maintain their independent levels
	if debugLogger.GetLevel() != zerolog.DebugLevel {
		t.Errorf("Debug logger should be at DEBUG level, got %v", debugLogger.GetLevel())
	}

	if errorLogger.GetLevel() != zerolog.ErrorLevel {
		t.Errorf("Error logger should be at ERROR level, got %v", errorLogger.GetLevel())
	}

	// Create a third logger at INFO level
	cfg3 := Config{
		LogDir: filepath.Join(tmpDir, "info"),
		Level:  "info",
	}
	infoLogger := New(cfg3)

	// Verify all three loggers still have their original levels
	if debugLogger.GetLevel() != zerolog.DebugLevel {
		t.Errorf("Debug logger level changed after creating other loggers, got %v", debugLogger.GetLevel())
	}

	if errorLogger.GetLevel() != zerolog.ErrorLevel {
		t.Errorf("Error logger level changed after creating other loggers, got %v", errorLogger.GetLevel())
	}

	if infoLogger.GetLevel() != zerolog.InfoLevel {
		t.Errorf("Info logger should be at INFO level, got %v", infoLogger.GetLevel())
	}
}
