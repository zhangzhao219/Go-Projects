// Code generated by hertz generator.

package main

import (
	handler "hertz_test/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.SdkController)

	// your code ...
}