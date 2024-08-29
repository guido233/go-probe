package bo

import "encoding/json"

/* 回复的统一格式 */

// Message 回复消息的通用格式
type Message struct {
	Id        int64           `json:"id"`
	Cmd       string          `json:"cmd"`
	Type      string          `json:"type"`
	Timestamp int64           `json:"timestamp"`
	Body      json.RawMessage `json:"body"`
}

/* 注册 */
type RegisterAckBody struct {
	Retcode    int    `json:"retcode"`
	RegisterId string `json:"registerid"`
}

/* 读取 */

// ReadReqBody 读取的请求消息
type ReadReqBody struct {
	Res   string          `json:"res"`
	Param json.RawMessage `json:"param"`
}

// ReadAckBody 读取的应答消息
type ReadAckBody struct {
	Retcode int   `json:"retcode"`
	ReadId  int64 `json:"readid"`
}

// 读取的上报消息
type ReadPutBody struct {
	Retcode int             `json:"retcode"`
	Readid  int64           `json:"readid"`
	Res     string          `json:"res"`
	Data    json.RawMessage `json:"data"`
}

/* 订阅 */

// ObserveReqBody 订阅请求消息
type ObserveReqBody struct {
	Res     string          `json:"res"`
	Gt      float64         `json:"gt"`      // 大于条件
	Lt      float64         `json:"lt"`      // 小于条件
	Tsample int             `json:"tsample"` // 采样频率
	Tmin    int             `json:"tmin"`    // 最小时间
	Tmax    int             `json:"tmax"`    // 最大时间
	Param   json.RawMessage `json:"param"`
}

// ObserveAckBody 订阅应答消息
type ObserveAckBody struct {
	Retcode   int   `json:"retcode"`
	ObserveId int64 `json:"observeid"`
}

// ObservePutBody 订阅的上报消息
type ObservePutBody struct {
	Retcode   int             `json:"retcode"`
	Observeid int64           `json:"observeid"`
	Res       string          `json:"res"`
	Data      json.RawMessage `json:"data"`
}

// ObserveCancelBody 订阅取消消息
type ObserveCancelBody struct {
	ObserveId []int64 `json:"observeid"`
}

/* 执行 */

// ExecReqBody 执行请求消息
type ExecReqBody struct {
	Res      string          `json:"res"`
	Args     json.RawMessage `json:"args"`
	Interval int             `json:"interval"` // 间隔执行时间，以毫秒为单位
	Times    int             `json:"times"`    // 执行次数
}

// ExecAckBody 执行应答消息
type ExecAckBody struct {
	Retcode int   `json:"code"`
	Execid  int64 `json:"execid"`
}

// ExecPutBody 执行的上报消息
type ExecPutBody struct {
	Retcode int             `json:"retcode"`
	Execid  int64           `json:"execid"`
	Res     string          `json:"res"`
	Data    json.RawMessage `json:"data"`
}

// ExecCancelBody 执行的取消消息
type ExecCancelBody struct {
	ExecId []int64 `json:"execid"`
	Force  bool    `json:"force,omitempty"` // 是否强制
}
