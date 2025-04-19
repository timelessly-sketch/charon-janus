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

			s.Group("/auth", func(group *ghttp.RouterGroup) {
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
			s.Group("/privilege", func(group *ghttp.RouterGroup) {
				group.Bind(
					privilege.Menu,
					privilege.Role,
				)
			})
			s.Run()
			return nil
		},
	}
)
