package input

import "charon-janus/internal/model/entity"

type MenuModelList struct {
	entity.AuthMenu
}

type MenuInput struct {
	entity.AuthMenu
}

type RoleModelList struct {
	entity.AuthRole
}

type RoleEditInput struct {
	PlatFormCode string `json:"platFormCode"`
	MenuIds      []int  `json:"menuIds"`
	ApiIds       []int  `json:"apiIds"`
	entity.AuthRole
}

type RoleDetailList struct {
	MenuList []MenuModelList `json:"menuList"`
	MenuIds  []int           `json:"menuIds"`
	ApiList  []ApiModelList  `json:"apiList"`
	ApiIds   []int           `json:"apiIds"`
}

type ApiModelList struct {
	entity.AuthApi
}

type ApiInput struct {
	Id           int    `json:"id"           description:"接口ID"`
	Pid          int    `json:"pid"          v:"required" description:"接口PID"`
	Name         string `json:"name"         v:"required" description:"接口名"`
	Icon         string `json:"icon"         v:"required" description:"图标"`
	Title        string `json:"title"        v:"required" description:"标题"`
	Path         string `json:"path"         v:"required-if:apiType,api" description:"接口路径"`
	Method       string `json:"method"       v:"required-if:apiType,api" description:"接口方法-目录为空,接口不能为空"`
	ApiType      string `json:"apiType"      v:"in:dir,api" description:"接口或者目录"`
	PlatformCode string `json:"platformCode" v:"required" description:"平台标识"`
}
