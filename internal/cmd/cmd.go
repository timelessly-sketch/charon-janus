package cmd

import (
	"charon-janus/internal/controller/login"
	"charon-janus/internal/controller/privilege"
	"charon-janus/internal/controller/system"
	"charon-janus/internal/library/cache"
	"charon-janus/internal/service"
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
			cache.SetAdapter(ctx)
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				ghttp.MiddlewareCORS,
				ghttp.MiddlewareHandlerResponse,
				service.Middleware().AuthMiddleware,
			}...)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					login.Login,
					system.PlatForm,
					system.User,
					privilege.Menu,
					privilege.Role,
					privilege.Api,
				)
			})
			s.Run()
			return nil
		},
	}
)
