package main

import (
	"path/filepath"

	"go.uber.org/zap"

	"github.com/FollowLille/daily-secretary-bot/internal/logger"
)

func main() {
	// Просто для тестирования логгирования
	configPath := filepath.Join("..", "..", "config", "logger.yaml")
	if err := logger.Init(configPath); err != nil {
		panic(err)
	}
	defer logger.Logger.Sync()

	logger.Logger.Info("Приложение запущено")
	logger.Logger.Debug("Отладочное сообщение")
	logger.Logger.Error("Тест ошибки")

	// Пример с дополнительными полями
	logger.Logger.Info("Пользователь авторизован",
		zap.String("username", "john_doe"),
		zap.Int("attempts", 3),
	)
}
