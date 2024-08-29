package main

import (
	"go-probe/conf"
	"go-probe/logger"
	"go-probe/serve"
	"go-probe/ws"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化日志
	defer logger.Start(
		logger.LogFilePath("./log/"),
		logger.LogSize(1),
		logger.LogMaxCount(3)).
		Stop()

	logger.Infof("go-probe start")

	// 初始化config
	conf.InitConfig()

	// 初始化ws
	ws.Init()
	defer ws.Close()

	// start serve
	serve.StartServe()

	// 保持存活
	c := initSignal()
	handleSignal(c)
}

// initSignal register signals handler.
func initSignal() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	return c
}

// handleSignal fetch signal from chan then do exit or reload.
func handleSignal(c chan os.Signal) {
	// Block until a signal is received.
	for {
		s := <-c
		logger.Infof("get a signal %s", s.String())
		switch s {
		case os.Interrupt:
			return
		case syscall.SIGHUP:
			//return
		default:
			return
		}
	}
}
