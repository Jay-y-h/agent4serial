package main

import (
	"agent/internal/config"
	"agent/internal/logger"
	"agent/internal/service"
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	
	configPath:=flag.String("config","configs/config.json","配置文件路径")
	flag.Parse()

	cfg,err:=config.Load(*configPath)
	if err!=nil{
		log.Fatalf("加载配置失败: %v", err)
	}
	if err:=logger.Init(cfg.Log);err!=nil{
		log.Fatalf("初始化日志失败: %v", err)
	}

	ctx,stop := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer stop()

	slog.Info("Serial Agent 启动", "gateway", cfg.Server.URL, "agent_id", cfg.Server.AgentID, "devices", len(cfg.Devices))

	if err := service.Run(ctx, cfg); err != nil {
		slog.Error("Serial Agent 启动失败", "error", err)
		os.Exit(1)
	}
	slog.Info("Serial Agent 已停止")
}