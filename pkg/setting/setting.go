package setting

import (
    "log"
    "time"

    "github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Redis struct {
    Host     string
    Password string
    DB       int
}
var RedisSetting = &Redis{}


var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
    var err error
    cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
    }

    mapTo("server", ServerSetting)
    mapTo("redis", RedisSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
    err := cfg.Section(section).MapTo(v)
    if err != nil {
        log.Fatalf("Cfg.MapTo %s err: %v", section, err)
    }
}