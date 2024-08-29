package test

import (
	"go-probe/logger"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	println(1 + 1)
}

func TestSplite(t *testing.T) {
	res := "DeviceInfo/#"
	// 获取字符串 某个字符之前的字符串
	retcodeStr := strings.Split(res, "/")

	if retcodeStr == nil {
		return
	}

	for _, r := range retcodeStr {
		logger.Infoln(r)
	}
}
