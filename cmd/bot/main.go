package bot

import (
	"log/slog"

	"github.com/Cladkoewka/grind-tracker/internal/config"
	"github.com/Cladkoewka/grind-tracker/internal/logger"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Log.LOG_LEVEL)
	slog.SetDefault(log)
	
}