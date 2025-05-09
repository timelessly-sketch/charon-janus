package middleware

import (
	"charon-janus/internal/library/token"
	"charon-janus/internal/model"
	"charon-janus/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

type sMiddleware struct{}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func (m *sMiddleware) AuthMiddleware(r *ghttp.Request) {
	var (
		handler      = r.GetServeHandler()
		method       = r.Method
		apiNotAuth   = g.Map{"code": http.StatusForbidden, "message": "接口未授权"}
		tokenMiss    = g.Map{"code": http.StatusUnauthorized, "message": "token缺失"}
		tokenInvalid = g.Map{"code": http.StatusUnauthorized, "message": "token解析异常"}
	)

	if handler.GetMetaTag("noAuth") == "true" || handler == nil {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatusExit(http.StatusUnauthorized, tokenMiss)
		return
	}
	tokenString := authHeader[len("Bearer "):]
	claims, err := token.ValidateJWT(tokenString)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnauthorized, tokenInvalid)
		return
	}

	flag, err := service.Api().AuthRoleApi(r.Context(), claims.Id, handler.Handler.Router.Uri, method)
	if err != nil {
		g.Log().Error(r.Context(), err.Error())
		r.Response.WriteStatusExit(http.StatusInternalServerError)
		return
	}
	if !flag {
		r.Response.WriteStatusExit(http.StatusOK, apiNotAuth)
		return
	}

	r.SetCtxVar("user", claims.Identity)
	r.Middleware.Next()
}

func (m *sMiddleware) GetUserIdentity(ctx context.Context) (user model.Identity) {
	if err := g.RequestFromCtx(ctx).GetCtxVar("user").Scan(&user); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}
