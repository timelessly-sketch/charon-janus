// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"context"
)

type (
	IPlatForm interface {
		List(ctx context.Context, inp *input.PageReq) (records []input.PlatFormModelList, total int, err error)
		Edit(ctx context.Context, inp *input.PlatFormEditInput) (err error)
		Options(ctx context.Context) (records []input.PlatFormModelList, err error)
		ProxyPath(ctx context.Context, path string, method string) (record entity.SysPlatform, err error)
	}
	IUser interface {
		List(ctx context.Context, inp *input.UserInput) (records []input.UserModelList, total int, err error)
		Detail(ctx context.Context, id int) (records input.UserModelDetail, err error)
		Edit(ctx context.Context, inp *input.UserEditInput) (err error)
		ResetPwd(ctx context.Context, inp *input.UserResetPwd) (err error)
	}
)

var (
	localPlatForm IPlatForm
	localUser     IUser
)

func PlatForm() IPlatForm {
	if localPlatForm == nil {
		panic("implement not found for interface IPlatForm, forgot register?")
	}
	return localPlatForm
}

func RegisterPlatForm(i IPlatForm) {
	localPlatForm = i
}

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
