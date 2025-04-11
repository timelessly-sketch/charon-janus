package cmd

import (
	"charon-janus/internal/controller/login"
	"charon-janus/internal/controller/system"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				ghttp.MiddlewareCORS,
				ghttp.MiddlewareHandlerResponse,
			}...)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					login.Login,
				)
			})
			s.Group("/system", func(group *ghttp.RouterGroup) {
				group.Bind(
					system.PlatForm,
					system.User,
				)
			})
			s.Run()
			return nil
		},
	}
)
