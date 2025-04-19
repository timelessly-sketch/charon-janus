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
	IMenu interface {
		List(ctx context.Context, code string) (records []input.MenuModelList, err error)
		Edit(ctx context.Context, inp *input.MenuInput) (err error)
	}
	IRole interface {
		List(ctx context.Context, code string) (records []input.RoleModelList, err error)
		Edit(ctx context.Context, inp *input.RoleEditInput) (err error)
		Detail(ctx context.Context, id int) (records input.RoleDetailList, err error)
	}
)

var (
	localMenu IMenu
	localRole IRole
)

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
