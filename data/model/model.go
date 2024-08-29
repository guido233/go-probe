package model

// Register 注册信息
type Register struct {
	DeviceId   string `json:"deviceid"`
	Authcode   string `json:"authcode"`
	Sdkmajor   int    `json:"sdk-major"`
	Sdkminor   int    `json:"sdk-minor"`
	Sdkmicro   int    `json:"sdk-micro"`
	Appversion string `json:"app-version"`
	Makedata   string `json:"make-data"`
}

// DeviceInfo Res: DeviceInfo/#
type DeviceInfo struct {
	DeviceID        string  `json:"DeviceID"`        // 设备唯一标识
	DeviceModel     string  `json:"DeviceModel"`     // 设备型号
	CpuModel        string  `json:"CpuModel"`        // CPU 型号
	HardwareVersion string  `json:"HardwareVersion"` // 设备硬件版本信息
	SoftwareVersion string  `json:"SoftwareVersion"` // 设备系统软件版本信息
	RamSize         float64 `json:"RamSize"`         // 内存大小，单位 MB
	FlashSize       float64 `json:"FlashSize"`       // FLASH 大小，单位 MB
	WifiAPCount     int32   `json:"WifiAPCount"`     // WiFi 模组数量
	LANCount        int32   `json:"LANCount"`        // LAN 口数量
	CellularCount   int32   `json:"CellularCount"`   // 蜂窝模组数量
	SerialPortCount int32   `json:"SerialPortCount"` // 串口数量
}

// ModuleCellular res: ModuleCellular/@IMEI/#
type ModuleCellular struct {
	Vendor  string `json:"Vendor"`
	Model   string `json:"Model"`
	Version string `json:"Version"`
	IMEI    string `json:"IMEI"`
	IMSI    string `json:"IMSI"`
}
