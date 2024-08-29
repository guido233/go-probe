package test

import "go-probe/logger"

func init() {
	// 初始化日志
	defer logger.Start(
		logger.LogFilePath("./log/"),
		logger.LogSize(1),
		logger.LogMaxCount(3)).
		Stop()

	logger.Infof("go-probe test")
}
