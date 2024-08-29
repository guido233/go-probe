package read

import (
	"encoding/json"
	"go-probe/constant"
	"go-probe/data/model/bo"
	"go-probe/logger"
	"go-probe/serve/res"
	"go-probe/ws"
	"time"
)

// 读取服务
func ReadServe(msg bo.Message) error {

	switch msg.Type {
	case constant.TypeAck:
		err := dealReadAck(msg.Body)
		if err != nil {
			logger.Errorf("read dealReadAck error: %v", err)
			return err
		}
	case constant.TypeReq:
		// 先响应
		sendReadAck(msg)
		// 处理请求
		err := dealReadReq(msg)
		if err != nil {
			logger.Errorf("read dealReadReq error: %v", err)
			return err
		}
	default:
		logger.Warnf("read msg.Type error: %v", msg.Type)
	}

	return nil
}

// 处理服务端的响应
func dealReadAck(data []byte) error {

	// 反序列化
	var ack bo.ReadAckBody
	err := json.Unmarshal(data, &ack)
	if err != nil {
		logger.Errorf("read json.Unmarshal error: %v", err)
		return err
	}

	logger.Infof("read ack retCode: %d ; readId: %s", ack.Retcode, ack.ReadId)

	return nil
}

// 向服务端发送响应
func sendReadAck(msg bo.Message) {

	m := bo.Message{
		Id:        msg.Id,
		Cmd:       msg.Cmd,
		Type:      constant.TypeAck,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	ack := bo.ReadAckBody{
		Retcode: 0,
		ReadId:  msg.Id,
	}

	body, _ := json.Marshal(ack)
	m.Body = body

	send, _ := json.Marshal(m)

	// 发送
	err := ws.Send(send)
	if err != nil {
		logger.Errorf("readId: %d sendReadAck error: %v", msg.Id, err)
	}

}

// 处理读取消息
func dealReadReq(msg bo.Message) error {

	// 反序列化
	var req bo.ReadReqBody
	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		logger.Errorf("read json.Unmarshal error: %v", err)
		return err
	}

	// 发送
	sendReadPut(msg, req.Res)

	return nil

}

func sendReadPut(msg bo.Message, r string) {

	// 获取资源的数据
	data := res.Res(r)

	m := bo.Message{
		Id:        msg.Id,
		Cmd:       msg.Cmd,
		Type:      constant.TypePut,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	put := bo.ReadPutBody{
		Retcode: 0,
		Readid:  msg.Id,
		Res:     r,
		Data:    data,
	}
	body, _ := json.Marshal(put)
	m.Body = body

	send, _ := json.Marshal(m)

	// 发送
	err := ws.Send(send)
	if err != nil {
		logger.Errorf("readId: %d sendReadPut error: %v", msg.Id, err)
	}

}
