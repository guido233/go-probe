package ws

import (
	"github.com/gorilla/websocket"
	"go-probe/conf"
	"go-probe/logger"
)

var Ws *websocket.Conn

func Init() {
	//服务器地址 websocket 统一使用 ws://
	url := "ws://" + conf.Conf.Websocket.Host + ":" + conf.Conf.Websocket.Port
	//使用默认拨号器，向服务器发送连接请求
	var err error
	Ws, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logger.Fatalf("ws connect error: %v", err)
	}
	logger.Infoln("ws connect success ! ! !")
}

func Close() {
	Ws.Close()
}

func Send(send []byte) error {

	logger.Infof("ws send: %s", send)

	err := Ws.WriteMessage(websocket.TextMessage, send)
	if err != nil {
		logger.Errorf("ws send error: %v", err)
		return err
	}

	return nil
}

func Read() ([]byte, error) {

	_, data, err := Ws.ReadMessage()
	if err != nil {
		logger.Errorf("ws.ReadMessage error: %v", err)
		return nil, err
	}

	logger.Infof("ws read: %s", data)

	return data, nil
}
