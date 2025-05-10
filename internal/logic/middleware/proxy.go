package middleware

import (
	"bytes"
	"charon-janus/internal/consts"
	"charon-janus/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func (m *sMiddleware) ProxyPlatform(r *ghttp.Request) {
	var (
		requestURI = r.URL.Path
		method     = r.Method
		ctx        = r.Context()
	)
	record, err := service.PlatForm().ProxyPath(ctx, m.cleanProxyPath(requestURI), method)
	if err != nil {
		g.Log().Warning(ctx, err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
		return
	}
	info := fmt.Sprintf("Api Path: %s, platform: %s", requestURI, record.PlatformCode)
	if record.PlatformCode == consts.App {
		g.Log().Infof(ctx, info)
		r.Middleware.Next()
		return
	}

	u, err := url.Parse(record.ServerUrl)
	if err != nil {
		g.Log().Errorf(ctx, "Invalid server URL: %s, err: %s", record.ServerUrl, err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
		return
	}

	proxyInfo := fmt.Sprintf("%s, source platform %s -> dest platform: %s", info, consts.App, record.PlatformCode)

	proxy := &httputil.ReverseProxy{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     10 * time.Second,
			DisableCompression:  true,
		},
		Director: func(req *http.Request) {
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
			req.URL.Path = r.URL.Path
			req.URL.RawQuery = r.URL.RawQuery
			req.Host = u.Host
			g.Log().Infof(ctx, "%s, server request: %s", proxyInfo, req.URL.String())
		},
		ErrorHandler: func(w http.ResponseWriter, req *http.Request, err error) {
			g.Log().Errorf(ctx, "%s, Err: %s", proxyInfo, err.Error())
			w.WriteHeader(http.StatusBadGateway)
		},

		ModifyResponse: func(resp *http.Response) error {
			body, _ := io.ReadAll(resp.Body)
			resp.Body = io.NopCloser(bytes.NewBuffer(body))

			g.Log().Infof(ctx, "%s, server response: %s", proxyInfo, body)

			for key, values := range resp.Header {
				r.Response.Header().Set(key, strings.Join(values, ","))
			}

			r.Response.WriteHeader(resp.StatusCode)

			return nil
		},
	}

	proxy.ServeHTTP(r.Response.RawWriter(), r.Request)
	r.ExitAll()
}
