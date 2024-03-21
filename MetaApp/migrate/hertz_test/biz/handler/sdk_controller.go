package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type SdkResponseErr struct {
	Return_code int    `json:"return_code"`
	Return_msg  string `json:"return_msg"`
}

func SdkController(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, SdkResponseErr{
		Return_code: 200,
		Return_msg:  "erty",
	})
}
