// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon-janus/internal/model/input"
	"context"
)

type (
	IApi interface {
		List(ctx context.Context, code string) (records []input.ApiModelList, err error)
		Edit(ctx context.Context, inp input.ApiInput) (err error)
		DetailApi(ctx context.Context, code string, id int) (records []input.ApiModelList, ids []int)
		RoleApiEdit(ctx context.Context, apiIds []int, roleId int) (err error)
		AuthRoleApi(ctx context.Context, userId int, path string, method string) (bool, error)
	}
	IMenu interface {
		List(ctx context.Context, code string) (records []input.MenuModelList, err error)
		Edit(ctx context.Context, inp *input.MenuInput) (err error)
		DetailMenu(ctx context.Context, code string, id int) (records []input.MenuModelList, ids []int)
		RoleMenuEdit(ctx context.Context, menuIds []int, roleId int) (err error)
	}
	IRole interface {
		List(ctx context.Context, code string) (records []input.RoleModelList, err error)
		Edit(ctx context.Context, inp *input.RoleEditInput) (err error)
		Detail(ctx context.Context, id int) (records input.RoleDetailList, err error)
	}
)

var (
	localApi  IApi
	localMenu IMenu
	localRole IRole
)

func Api() IApi {
	if localApi == nil {
		panic("implement not found for interface IApi, forgot register?")
	}
	return localApi
}

func RegisterApi(i IApi) {
	localApi = i
}

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
