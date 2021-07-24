package setting

import (
	"github.com/go-ini/ini"
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

func init() {
	var err error
	Cfg, err = ini.Load("../../conf/api.ini")
	if err != nil {
		log.Fatalf("配置文件载入失败: %v",err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer(){
	sec,err := Cfg.GetSection("server")

	if err != nil {
		log.Fatalf("配置文件格式错误：没有server")
	}
	HTTPPort = sec.Key("PORT").MustInt(5000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60))  * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp(){
	sec,err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("配置文件格式错误：没有app")
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}