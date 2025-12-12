package logger

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger wraps zerolog.Logger with additional functionality
type Logger struct {
	zerolog.Logger
	fileWriter io.Closer // Reference to lumberjack writer for proper cleanup
}

// Config holds logger configuration
type Config struct {
	Level         string      // debug, info, warn, error
	LogDir        string
	Filename      string      // Log filename (default: "go.log")
	MaxSizeMB     int
	MaxBackups    int
	Console       bool        // Enable console output
	DirMode       os.FileMode // Directory permissions (default: 0750)
	DisableCaller bool        // Disable caller info (file:line) in logs for privacy (default: false/enabled)
}

// New creates a new logger instance
func New(cfg Config) *Logger {
	// Set defaults
	if cfg.LogDir == "" {
		cfg.LogDir = "./logs"
	}
	if cfg.Filename == "" {
		cfg.Filename = "go.log"
	}
	if cfg.MaxSizeMB == 0 {
		cfg.MaxSizeMB = 10
	}
	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = 5
	}
	if cfg.DirMode == 0 {
		cfg.DirMode = 0750 // rwxr-x--- (more secure default)
	}

	// Sanitize and validate paths to prevent path traversal attacks
	cfg.LogDir = filepath.Clean(cfg.LogDir)
	cfg.Filename = filepath.Clean(cfg.Filename)

	// Check for path traversal attempts in LogDir
	if strings.Contains(cfg.LogDir, "..") {
		// Path traversal detected - fall back to stderr with warning
		return createStderrLogger("path traversal detected in LogDir: " + cfg.LogDir)
	}

	// Validate filename doesn't contain path separators
	if strings.ContainsAny(cfg.Filename, `/\`) || strings.Contains(cfg.Filename, "..") {
		// Invalid filename - fall back to stderr with warning
		return createStderrLogger("invalid filename (contains path separators or traversal): " + cfg.Filename)
	}

	// Create log directory if it doesn't exist
	if err := os.MkdirAll(cfg.LogDir, cfg.DirMode); err != nil {
		// Log the error to stderr using structured logging before falling back
		stderrLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		stderrLogger.Error().
			Err(err).
			Str("log_dir", cfg.LogDir).
			Msg("Failed to create log directory, falling back to stderr")

		return &Logger{Logger: stderrLogger}
	}

	// Parse log level (set per-logger, not globally)
	level := parseLogLevel(cfg.Level)

	// Configure file rotation
	fileWriter := &lumberjack.Logger{
		Filename:   filepath.Join(cfg.LogDir, cfg.Filename),
		MaxSize:    cfg.MaxSizeMB,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     30, // days
		Compress:   false,
	}

	// Create multi-writer (file + console if enabled)
	var writers []io.Writer
	writers = append(writers, fileWriter)

	if cfg.Console {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
			NoColor:    false,
		}
		writers = append(writers, consoleWriter)
	}

	multiWriter := io.MultiWriter(writers...)

	// Create logger context with per-instance level (not global)
	ctx := zerolog.New(multiWriter).
		Level(level).
		With().
		Timestamp()

	// Conditionally add caller information based on configuration
	// By default (DisableCaller = false), caller info is included for debugging
	// Set DisableCaller = true to omit file paths for enhanced privacy/security
	if !cfg.DisableCaller {
		ctx = ctx.Caller()
	}

	logger := ctx.Logger()

	return &Logger{
		Logger:     logger,
		fileWriter: fileWriter, // Store for proper cleanup on Close()
	}
}

// createStderrLogger creates a logger that writes to stderr with a security warning
func createStderrLogger(warningMsg string) *Logger {
	// Log security warning to stderr
	stderrLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	stderrLogger.Error().
		Str("security_warning", warningMsg).
		Msg("SECURITY: Invalid logger configuration, falling back to stderr")

	return &Logger{Logger: stderrLogger}
}

// parseLogLevel converts string log level to zerolog level
func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn", "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}

// Close closes the logger and flushes any buffered logs
func (l *Logger) Close() error {
	// Close the underlying file writer to ensure all logs are flushed
	if l.fileWriter != nil {
		return l.fileWriter.Close()
	}
	return nil
}

// WithField adds a field to the logger
func (l *Logger) WithField(key string, value interface{}) *Logger {
	newLogger := l.Logger.With().Interface(key, value).Logger()
	return &Logger{
		Logger:     newLogger,
		fileWriter: l.fileWriter, // Preserve fileWriter reference
	}
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	ctx := l.Logger.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}
	newLogger := ctx.Logger()
	return &Logger{
		Logger:     newLogger,
		fileWriter: l.fileWriter, // Preserve fileWriter reference
	}
}

// WithError adds an error to the logger context
func (l *Logger) WithError(err error) *Logger {
	newLogger := l.Logger.With().Err(err).Logger()
	return &Logger{
		Logger:     newLogger,
		fileWriter: l.fileWriter, // Preserve fileWriter reference
	}
}
