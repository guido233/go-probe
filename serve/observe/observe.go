package observe

import (
	"encoding/json"
	"go-probe/constant"
	"go-probe/data/model/bo"
	"go-probe/logger"
	"go-probe/serve/res"
	"go-probe/ws"
	"time"
)

// 定义一个通道
var observer = make(map[int64]*Observer)

type Observer struct {
	ObserveId int64
	C         chan struct{}
}

func newObserver(observeId int64) *Observer {
	return &Observer{
		ObserveId: observeId,
		C:         make(chan struct{}),
	}
}

// ObserveServe 订阅服务
func ObserveServe(msg bo.Message) error {

	switch msg.Type {
	case constant.TypeAck:
		err := dealObserveAck(msg.Body)
		if err != nil {
			return err
		}
	case constant.TypeReq:
		// 先响应
		sendObserveAck(msg)
		// 处理请求
		go dealObserveReq(msg)
	case constant.TypeCancel:
		err := dealObserveCancel(msg)
		if err != nil {
			return err
		}
	default:
		logger.Warnf("observe msg.Type error: %v", msg.Type)
	}

	return nil
}

// 处理服务端的响应
func dealObserveAck(data []byte) error {

	var ack bo.ObserveAckBody
	err := json.Unmarshal(data, &ack)
	if err != nil {
		logger.Errorf("observe json.Unmarshal error: %v", err)
		return err
	}

	logger.Infof("observe ack retCode: %d ; observeId: %s", ack.Retcode, ack.ObserveId)

	return nil
}

// 向服务端发送响应
func sendObserveAck(msg bo.Message) {

	m := bo.Message{
		Id:        msg.Id,
		Cmd:       msg.Cmd,
		Type:      constant.TypeAck,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	ack := bo.ObserveAckBody{
		Retcode:   0,
		ObserveId: msg.Id,
	}

	body, _ := json.Marshal(ack)
	m.Body = body

	send, _ := json.Marshal(m)

	// 发送
	err := ws.Send(send)
	if err != nil {
		logger.Errorf("observeId: %d sendObserveAck error: %v", msg.Id, err)
	}

}

func dealObserveReq(msg bo.Message) {

	ob := newObserver(msg.Id)
	observer[msg.Id] = ob

	// 反序列化
	var req bo.ObserveReqBody
	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		logger.Errorf("observe json.Unmarshal error: %v", err)
		return
	}

	for {
		select {
		case <-ob.C:
			return
		default:
			// 发送
			sendObservePut(msg, req.Res)
			if req.Tsample != 0 {
				time.Sleep(time.Duration(req.Tsample) * time.Millisecond)
			}
		}
	}

}

func sendObservePut(msg bo.Message, r string) {

	// 获取资源的数据
	data := res.Res(r)

	m := bo.Message{
		Id:        msg.Id,
		Cmd:       msg.Cmd,
		Type:      constant.TypePut,
		Timestamp: time.Now().Unix(),
		Body:      nil,
	}

	put := bo.ObservePutBody{
		Retcode:   0,
		Observeid: msg.Id,
		Res:       r,
		Data:      data,
	}
	body, _ := json.Marshal(put)
	m.Body = body

	send, _ := json.Marshal(m)

	// 发送
	err := ws.Send(send)
	if err != nil {
		logger.Errorf("observeId: %d sendObservePut error: %v", msg.Id, err)
	}

}

// 处理订阅取消
func dealObserveCancel(msg bo.Message) error {

	var ack bo.ObserveCancelBody
	err := json.Unmarshal(msg.Body, &ack)
	if err != nil {
		logger.Errorf("observe cancel json.Unmarshal error: %v", err)
		return err
	}

	logger.Infof("observe cancel observeId: %v", ack.ObserveId)

	// 循环要取消的订阅
	for _, r := range ack.ObserveId {
		// map中查找是否有这个id
		if _, ok := observer[r]; ok {
			// 关闭通道
			close(observer[r].C)
			// 删掉这个订阅
			delete(observer, r)
		}
	}

	// 响应订阅取消
	sendObserveAck(msg)

	return nil
}
