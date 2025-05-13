package contexts

import (
	"charon-janus/internal/consts"
	"charon-janus/internal/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextHTTPKey, customCtx)
}

func Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextHTTPKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func SetResponse(ctx context.Context, response *ghttp.DefaultHandlerResponse) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetResponse, c == nil ")
		return
	}
	c.Response = response
}

func SetUser(ctx context.Context, user *model.Identity) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetUser, c == nil ")
		return
	}
	c.User = user
}
