package serve

import (
	"encoding/json"
	"go-probe/constant"
	"go-probe/data/model/bo"
	"go-probe/logger"
	"go-probe/serve/execute"
	"go-probe/serve/observe"
	"go-probe/serve/read"
	"go-probe/serve/register"
	"go-probe/ws"
)

func StartServe() {

	// 注册设备
	err := register.Register()
	if err != nil {
		logger.Errorf("register error: %v", err)
	}

	// 接收消息
	for {

		logger.Infoln("开始接收消息")

		r, err := ws.Read()
		if err != nil {
			logger.Errorf("serve ws.Read error: %v", err)
		}

		// 反序列化
		var msg bo.Message
		err = json.Unmarshal(r, &msg)
		if err != nil {
			logger.Errorf("serve json.Unmarshal error: %v", err)
		}

		switch msg.Cmd {
		case constant.CmdRegister: /* 注册 */
			err := register.RegisterServe(msg)
			if err != nil {
				logger.Errorf("register error: %v", err)
			}
		case constant.CmdRead: /* 读取 */
			err := read.ReadServe(msg)
			if err != nil {
				logger.Errorf("read error: %v", err)
			}
		case constant.CmdObserve: /* 订阅 */
			err := observe.ObserveServe(msg)
			if err != nil {
				logger.Errorf("observe error: %v", err)
			}
		case constant.CmdExecute: /* 执行 */
			err := execute.ExecServer(msg)
			if err != nil {
				logger.Errorf("execute error: %v", err)
			}
		case constant.CmdReport: /* 上报 */
		case constant.CmdWrite: /* 写 */
		default:
			if msg.Cmd != "" {
				logger.Warnf("serve cmd error: %v", msg.Cmd)
			}
		}

		//time.Sleep(time.Second * 1)
	}

}
