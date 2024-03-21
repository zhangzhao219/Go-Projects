package router

import (
	"github.com/beego/beego/v2/server/web"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type SdkController struct {
	web.Controller
}

func (c *SdkController) Post() {

	c.Data["json"] = &Response{
		Code:    "200",
		Message: "pong",
	}
	_ = c.ServeJSON()
}
