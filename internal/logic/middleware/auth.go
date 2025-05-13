package middleware

import (
	"charon-janus/internal/library/contexts"
	"charon-janus/internal/library/token"
	"charon-janus/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

func (m *sMiddleware) AuthMiddleware(r *ghttp.Request) {
	var (
		handler      = r.GetServeHandler()
		method       = r.Method
		ctx          = r.Context()
		apiNotAuth   = g.Map{"code": http.StatusForbidden, "message": "接口未授权"}
		tokenMiss    = g.Map{"code": http.StatusUnauthorized, "message": "token缺失"}
		tokenInvalid = g.Map{"code": http.StatusUnauthorized, "message": "token解析异常"}
	)
	if handler.GetMetaTag("noAuth") == "true" {
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

	flag, err := service.Api().AuthRoleApi(ctx, claims.Id, m.cleanProxyPath(r.URL.Path), method)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		r.Response.WriteStatusExit(http.StatusInternalServerError)
		return
	}
	if !flag {
		r.Response.WriteStatusExit(http.StatusOK, apiNotAuth)
		return
	}
	contexts.SetUser(ctx, &claims.Identity)
	r.Middleware.Next()
}

func (m *sMiddleware) cleanProxyPath(path string) string {
	if path == "" || path == "/" {
		return "/"
	}

	parts := gstr.Split(path, "/")
	var cleanParts []string

	for _, p := range parts {
		if p != "" {
			cleanParts = append(cleanParts, p)
		}
	}

	if len(cleanParts) > 0 && gregex.IsMatchString(`^\d+$`, cleanParts[len(cleanParts)-1]) {
		cleanParts = cleanParts[:len(cleanParts)-1]
	}

	return "/" + gstr.Join(cleanParts, "/")
}
