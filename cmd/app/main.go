package main

import (
	"github.com/k15zzz/html2pdf_on_go/pkg/logger"
	"github.com/k15zzz/html2pdf_on_go/pkg/utils"
	"log"
	"os"

	"github.com/k15zzz/html2pdf_on_go/configs"
	"github.com/k15zzz/html2pdf_on_go/internal/server"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := configs.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	s := server.NewServer(cfg, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
