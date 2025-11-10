package logservice

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	// logDir определяет имя директории, где будут храниться логи игр.
	logDir = "log"
	// logFile определяет имя файла для записи истории игр.
	logFile = "game_history.log"
)

// FullLogger реализует интерфейс IFullLogger.
// Он предоставляет функционал для логирования системных событий в консоль
// и записи истории игр в файл.
type FullLogger struct{}

// NewFullLogger создает и возвращает экземпляр FullLogger как IFullLogger.
// При инициализации настраивает формат вывода системного лога.
func NewFullLogger() IFullLogger {
	// Настраиваем log-пакет для вывода системных сообщений,
	// чтобы он включал дату/время и имя файла
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return &FullLogger{}
}

// LogGame записывает информацию об одном раунде игры в файл game_history.log.
// Реализует метод интерфейса IGameLogger.
func (f *FullLogger) LogGame(name string, bet float64, result string, balance float64) error {
	// 1. Создаем папку 'log/', если она не существует
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("не удалось создать директорию %s: %w", logDir, err)
	}

	fullPath := filepath.Join(logDir, logFile)

	// 2. Открываем файл для добавления (append) или создаем его
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл лога %s: %w", fullPath, err)
	}

	defer file.Close()

	// 3. Форматируем и записываем строку лога
	logEntry := fmt.Sprintf("[%s] | Name: %s | Bet: %.2f | Result: %s | New Balance: %.2f\n",
		time.Now().Format("2006-01-02 15:04:05"),
		name,
		bet,
		result,
		balance)

	if _, err := file.WriteString(logEntry); err != nil {
		return fmt.Errorf("не удалось записать в файл лога: %w", err)
	}
	return nil
}

// Info выводит информационное сообщение с тегом [INFO] в консоль.
// Реализует метод интерфейса ISystemLogger
func (c *FullLogger) Info(format string, a ...interface{}) {
	log.Printf("[INFO] "+format, a...)
}

// Warn выводит предупреждение с тегом [WARN] в консоль.
// Реализует метод интерфейса ISystemLogger.
func (c *FullLogger) Warn(format string, a ...interface{}) {
	log.Printf("[WARN] "+format, a...)
}

// Error выводит сообщение об ошибке с тегом [ERROR] в консоль.
// Реализует метод интерфейса ISystemLogger.
func (c *FullLogger) Error(format string, a ...interface{}) {
	log.Printf("[ERROR] "+format, a...)
}

// Fatal выводит критическую ошибку с тегом [FATAL] и завершает работу приложения (os.Exit(1)).
// Реализует метод интерфейса ISystemLogger.
func (c *FullLogger) Fatal(format string, a ...interface{}) {
	log.Printf("[FATAL] "+format, a...)
}
