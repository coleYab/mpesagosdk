package mpesasdk

import (
	"github.com/coleYab/mpesasdk/config"
)

type App struct {
	cfg *config.Config
}

// Creates a new mpesa App
func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}
