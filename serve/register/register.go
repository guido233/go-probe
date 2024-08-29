package register

import (
	"encoding/json"
	"go-probe/constant"
	"go-probe/data/model"
	"go-probe/data/model/bo"
	"go-probe/logger"
	"go-probe/ws"
	"time"
)

func Register() error {
	req := bo.Message{
		Id:        1,
		Cmd:       constant.CmdRegister,
		Type:      constant.TypeReq,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	register := model.Register{
		DeviceId:   "987654321",
		Authcode:   "12345",
		Sdkmajor:   1,
		Sdkminor:   1,
		Sdkmicro:   1,
		Appversion: "1.0.0",
		Makedata:   time.Now().Format("2006-01-02 15:04:05"),
	}

	// 序列化
	body, err := json.Marshal(register)
	if err != nil {
		logger.Errorf("register json.Marshal error: %v", err)
		return err
	}

	req.Body = body

	send, err := json.Marshal(req)
	if err != nil {
		logger.Errorf("register req json.Marshal error: %v", err)
		return err
	}

	err = ws.Send(send)
	if err != nil {
		logger.Errorf("register ws.Send error: %v", err)
		return err
	}

	return nil
}

func RegisterServe(msg bo.Message) error {

	switch msg.Type {
	case constant.TypeAck:
		err := dealRegisterAck(msg.Body)
		if err != nil {
			logger.Errorf("register dealRegisterAck error: %v", err)
			return err
		}
	default:
		logger.Warnf("register msg.Type error: %v", msg.Type)
	}

	return nil
}

func dealRegisterAck(data []byte) error {

	// 反序列化
	var ack bo.RegisterAckBody
	err := json.Unmarshal(data, &ack)
	if err != nil {
		logger.Errorf("register json.Unmarshal error: %v", err)
		return err
	}

	logger.Infof("register ack retCode: %d ; registerId: %s", ack.Retcode, ack.RegisterId)

	return nil
}
