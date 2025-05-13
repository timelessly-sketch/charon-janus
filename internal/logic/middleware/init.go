package middleware

import (
	"charon-janus/internal/library/contexts"
	"charon-janus/internal/model"
	"charon-janus/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func (m *sMiddleware) Init(r *ghttp.Request) {
	var (
		path   = m.cleanProxyPath(r.URL.Path)
		method = r.Method
	)

	data := &model.HandlerRequest{
		Path:   path,
		Body:   r.GetBodyString(),
		Method: method,
	}

	contexts.Init(r, &model.Context{
		Request: data,
	})
	r.Middleware.Next()
}
