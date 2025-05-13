// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure of table sys_log for DAO operations like Where/Data.
type SysLog struct {
	g.Meta     `orm:"table:sys_log, do:true"`
	Id         interface{} //
	ReqId      interface{} // 日志id
	Username   interface{} // 用户
	Url        interface{} // api路径
	Method     interface{} // 方法
	GetData    *gjson.Json // get数据
	PostData   *gjson.Json // post数据
	HeaderData *gjson.Json // 请求头
	ErrorData  *gjson.Json // 报错数据
	UserAgent  interface{} // UA信息
	Browser    interface{} // 浏览器
	TakeUpTime interface{} // 请求耗时
	ClientIp   interface{} //
	Timestamp  *gtime.Time // 响应时间
	Code       interface{} // 响应码
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 创建时间
}
