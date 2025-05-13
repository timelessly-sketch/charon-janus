package location

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func GetClientIp(r *ghttp.Request) string {
	if r == nil {
		return ""
	}
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}

	if gstr.Contains(ip, ",") {
		ip = gstr.StrTillEx(ip, ",")
	}

	if gstr.Contains(ip, ", ") {
		ip = gstr.StrTillEx(ip, ", ")
	}
	return ip
}
