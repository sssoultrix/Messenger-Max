package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger(env string) {
	var handler slog.Handler

	switch env {
	case "dev":
		// Красивый читаемый вывод в консоль
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	default:
		// JSON — удобно собирать в ELK/Loki
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	Log = slog.New(handler)
}
