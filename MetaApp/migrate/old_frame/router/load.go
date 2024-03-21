package router

import (
	"old_frame/config"

	"github.com/beego/beego/v2/server/web"
)

func Load(serverConfig *config.ServerConfig) {
	web.BConfig.CopyRequestBody = true
	web.BConfig.ServerName = serverConfig.ServiceConf.HTTPServiceName
	web.BConfig.Listen.HTTPAddr = serverConfig.ListenConf.HttpAddr
	web.BConfig.Listen.HTTPPort = serverConfig.ListenConf.HttpPort

	web.Router("/ping", &SdkController{})
}
