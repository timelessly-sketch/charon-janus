// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure for table sys_log.
type SysLog struct {
	Id         int         `json:"id"         orm:"id"           description:""`       //
	ReqId      string      `json:"reqId"      orm:"req_id"       description:"日志id"`   // 日志id
	Username   string      `json:"username"   orm:"username"     description:"用户"`     // 用户
	Url        string      `json:"url"        orm:"url"          description:"api路径"`  // api路径
	Method     string      `json:"method"     orm:"method"       description:"方法"`     // 方法
	GetData    *gjson.Json `json:"getData"    orm:"get_data"     description:"get数据"`  // get数据
	PostData   *gjson.Json `json:"postData"   orm:"post_data"    description:"post数据"` // post数据
	HeaderData *gjson.Json `json:"headerData" orm:"header_data"  description:"请求头"`    // 请求头
	ErrorData  *gjson.Json `json:"errorData"  orm:"error_data"   description:"报错数据"`   // 报错数据
	UserAgent  string      `json:"userAgent"  orm:"user_agent"   description:"UA信息"`   // UA信息
	Browser    string      `json:"browser"    orm:"browser"      description:"浏览器"`    // 浏览器
	TakeUpTime string      `json:"takeUpTime" orm:"take_up_time" description:"请求耗时"`   // 请求耗时
	ClientIp   string      `json:"clientIp"   orm:"client_ip"    description:""`       //
	Timestamp  *gtime.Time `json:"timestamp"  orm:"timestamp"    description:"响应时间"`   // 响应时间
	Code       int         `json:"code"       orm:"code"         description:"响应码"`    // 响应码
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`   // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:"创建时间"`   // 创建时间
}
