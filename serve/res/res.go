package res

import (
	"bytes"
	"go-probe/constant"
	"go-probe/logger"
	"os/exec"
)

func Res(res string) []byte {

	var data []byte

	// 获取字符串 某个字符之前的字符串
	//resStr := strings.Split(res, "/")[0]

	switch res {
	case constant.DeviceInfoRes:
		data = bashCmd(constant.DeviceInfoBash)
	default:
		logger.Warnf("Res res error: %v", res)
	}

	if data == nil {
		logger.Errorf("res data is nil")
		return nil
	}

	return data
}

func bashCmd(sh string) []byte {

	// 创建一个新的命令
	cmd := exec.Command("sudo", "bash", "sh/"+sh)
	// 创建一个用于保存命令输出的缓冲区
	var out bytes.Buffer
	// 将命令的输出重定向到缓冲区
	cmd.Stdout = &out
	// 运行命令
	err := cmd.Run()
	if err != nil {
		logger.Errorf("%s Failed to execute command: %v", sh, err)
		return nil
	}

	return out.Bytes()
}
