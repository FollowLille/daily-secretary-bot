package logger

import (
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once   sync.Once
	Logger *CompositeLogger
)

type CompositeLogger struct {
	commonLog *zap.Logger
	errorLog  *zap.Logger
	outputLog *zap.Logger
	console   *zap.Logger
}

func Init(configPath string) error {
	var initErr error

	once.Do(func() {
		cfg, err := LoadConfig(configPath)
		if err != nil {
			initErr = err
			return
		}

		// Файловые логи
		commonCore := createFileCore(cfg.LogDir, cfg.CommonLog, zapcore.InfoLevel)
		errorCore := createFileCore(cfg.LogDir, cfg.ErrorLog, zapcore.ErrorLevel)
		outputCore := createFileCore(cfg.LogDir, cfg.OutputLog, zapcore.DebugLevel)

		// Консольный лог
		consoleLevel := zapcore.DebugLevel
		switch cfg.Console.Level {
		case "info":
			consoleLevel = zapcore.InfoLevel
		case "error":
			consoleLevel = zapcore.ErrorLevel
		}

		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(consoleLevel),
		)

		Logger = &CompositeLogger{
			commonLog: zap.New(commonCore),
			errorLog:  zap.New(errorCore),
			outputLog: zap.New(outputCore),
			console:   zap.New(consoleCore).WithOptions(zap.AddCaller()),
		}
	})

	return initErr
}

func createFileCore(logDir string, cfg LogFileConfig, level zapcore.Level) zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   filepath.Join(logDir, cfg.Filename),
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}),
		zap.NewAtomicLevelAt(level),
	)
}
