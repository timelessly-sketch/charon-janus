package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Identity struct {
	Id       int    `json:"id" dc:"id"`
	Nickname string `json:"nickname" dc:"昵称"`
	Username string `json:"username" dc:"英文名"`
	Name     string `json:"name" dc:"中文名"`
	UserId   string `json:"userId" dc:"userId"`
}

type HandlerRequest struct {
	Path   string `json:"path" dc:"path"`
	Body   string `json:"body" dc:"body"`
	Method string `json:"method" dc:"method"`
}

type Context struct {
	User     *Identity
	Response *ghttp.DefaultHandlerResponse
	Request  *HandlerRequest
}
