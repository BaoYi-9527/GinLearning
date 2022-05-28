// Package setting
// 键里调用 /conf/app.ini 文件下配置的模块
package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSecret string
)

func init()  {
	var err error
	// 加载配置文件
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil{
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	// 加载相关配置
	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase 加载应用运行模式
func LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 加载 server 分区
func LoadServer()  {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTPPort").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	ReadTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp 加载 app 分区
func LoadApp()  {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize  = sec.Key("PAGE_SIZE").MustInt(10)
}


