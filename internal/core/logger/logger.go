package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey struct{}

var (
	loggerKey = contextKey{}
)

type Logger struct {
	*zap.Logger
	file *os.File
}

func (log *Logger) Close() {
	if err := log.file.Close(); err != nil {
		fmt.Printf("failed to close application logger: %v\n", err)
	}
}

func NewLogger(config Config) (*Logger, error) {
	consoleLevel, err := zap.ParseAtomicLevel(config.Level)
	if err != nil {
		return nil, fmt.Errorf("parse atomic logger level: %w", err)
	}

	err = os.MkdirAll(config.Dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("make log dir: %w", err)
	}
	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	path := filepath.Join(
		config.Dir,
		fmt.Sprintf("%s.log", timestamp),
	)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	cfgEncoder := zap.NewDevelopmentEncoderConfig()
	cfgEncoder.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")
	consoleEncoder := zapcore.NewConsoleEncoder(cfgEncoder)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), consoleLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(file), zapcore.InfoLevel),
	)
	logger := zap.New(core)
	return &Logger{
		Logger: logger,
		file:   file,
	}, nil
}

func (log *Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

func FromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		panic("no logger in context")
	}

	return log
}
