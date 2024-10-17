package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CreateLogger() (*zap.SugaredLogger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Capitalize the log level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC timestamp format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Short caller (file and line)
	}

	// Set up lumberjack as a logger:
	logger := &lumberjack.Logger{
		Filename:   "./logs/dsm.log", // Or any other path
		MaxSize:    500,              // MB; after this size, a new log file is created
		MaxBackups: 3,                // Number of backups to keep
		MaxAge:     28,               // Days
		Compress:   true,             // Compress the backups using gzip
	}

	writeSyncer := zapcore.AddSync(logger)

	// Set up zap logger configuration:
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Using JSON encoder, but you can choose another
		writeSyncer,
		zapcore.DebugLevel,
	)

	loggerZap := zap.New(core)
	sugar := loggerZap.Sugar()
	defer func(sugar *zap.SugaredLogger) {
		err := sugar.Sync()
		if err != nil {
			return
		}
	}(sugar)

	return sugar, nil
}
