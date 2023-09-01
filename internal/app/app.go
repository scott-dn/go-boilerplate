package app

import (
	"github.com/scott-dn/go-boilerplate/configs"
	"github.com/scott-dn/go-boilerplate/internal/pkg/logger"
)

type App struct {
}

func Init() *App {
	config := configs.NewConfig()
	logger.InitGlobal(config)

	return &App{}
}
