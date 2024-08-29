package execute

import (
	"encoding/json"
	"go-probe/constant"
	"go-probe/data/model/bo"
	"go-probe/logger"
	"go-probe/serve/res"
	"go-probe/ws"
	"time"
)

var executer = make(map[int64]*Executer)

type Executer struct {
	ExecId int64
	C      chan struct{}
}

func newExecuter(execid int64) *Executer {
	return &Executer{
		ExecId: execid,
		C:      make(chan struct{}),
	}
}

// ExecServer
func ExecServer(msg bo.Message) error {
	switch msg.Type {
	case constant.TypeAck:
		err := dealExecAck(msg.Body)
		if err != nil {
			return err
		}
	case constant.TypeReq:
		// 先响应
		sendExecAck(msg)
		// 处理请求
		go dealExecReq(msg)
	case constant.TypeCancel:
		err := dealExecCancel(msg)
		if err != nil {
			return err
		}
	default:
		logger.Warnf("exec msg.Type error: %v", msg.Type)
	}
	return nil
}

// 处理执行响应
func dealExecAck(data []byte) error {

	var ack bo.ExecAckBody
	err := json.Unmarshal(data, &ack)
	if err != nil {
		logger.Errorf("exec json.Unmarshal error: %v", err)
		return err
	}

	logger.Errorf("exec ack retCode: %v; execId: %s", ack.Retcode, ack.Execid)

	return nil
}

// 向服务端发送响应
func sendExecAck(msg bo.Message) {

	m := bo.Message{
		Id:        msg.Id,
		Cmd:       msg.Cmd,
		Type:      constant.TypeAck,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	ack := bo.ExecAckBody{
		Retcode: 0,
		Execid:  msg.Id,
	}

	body, _ := json.Marshal(ack)
	m.Body = body

	send, _ := json.Marshal(m)

	// 发送
	err := ws.Send(send)
	if err != nil {
		logger.Errorf("execId: %d sendExecAck error: %v", msg.Id, err)
	}

}

func dealExecReq(msg bo.Message) {

	ob := newExecuter(msg.Id)
	executer[msg.Id] = ob

	// 反序列化
	var req bo.ExecReqBody
	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		logger.Errorf("exec json.Unmarshal error: %v", err)
		return
	}
	//times := req.Times

	for {
		select {
		case <-ob.C:
			return
		default:
			// 发送
			sendExecPut(msg, req.Res)
			if req.Interval != 0 {
				time.Sleep(time.Duration(req.Interval) * time.Millisecond)
			}
		}
	}

}

func sendExecPut(msg bo.Message, r string) {

	// 获取资源的数据
	data := res.Res(r)

	m := bo.Message{
		Id:        msg.Id,
		Cmd:       msg.Cmd,
		Type:      constant.TypePut,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	put := bo.ExecPutBody{
		Retcode: 0,
		Execid:  msg.Id,
		Res:     r,
		Data:    data,
	}
	body, _ := json.Marshal(put)
	m.Body = body

	send, _ := json.Marshal(m)

	// 发送
	err := ws.Send(send)
	if err != nil {
		logger.Errorf("ExecId: %d sendExecPut error: %v", msg.Id, err)
	}

}

// 处理订阅取消
func dealExecCancel(msg bo.Message) error {

	var ack bo.ExecCancelBody
	err := json.Unmarshal(msg.Body, &ack)
	if err != nil {
		logger.Errorf("exec cancel json.Unmarshal error: %v", err)
		return err
	}

	logger.Infof("exec cancel execId: %v", ack.ExecId)

	// 循环要取消的订阅
	for _, r := range ack.ExecId {
		// map中查找是否有这个id
		if _, ok := executer[r]; ok {
			// 关闭通道
			close(executer[r].C)
			// 删掉这个订阅
			delete(executer, r)
		}
	}

	// 响应订阅取消
	sendExecAck(msg)

	return nil
}
