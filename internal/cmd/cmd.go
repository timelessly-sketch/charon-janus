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

			s.BindHookHandler("/*any", ghttp.HookAfterOutput, service.Hook().AfterOutput)

			s.BindMiddlewareDefault([]ghttp.HandlerFunc{
				ghttp.MiddlewareCORS,
				service.Middleware().Init,
				service.Middleware().AuthMiddleware,
				service.Middleware().ProxyPlatform,
				service.Middleware().ResponseHandler,
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
			return
		},
	}
)
