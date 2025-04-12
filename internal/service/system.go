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
	ILogin interface {
		Login(ctx context.Context, inp *input.AccountLoginInp) (records input.LoginModel, err error)
	}
	IPlatForm interface {
		List(ctx context.Context, inp *input.PageReq) (records []input.PlatFormModelList, total int, err error)
		Edit(ctx context.Context, inp *input.PlatFormEditInput) (err error)
		Options(ctx context.Context) (records []input.PlatFormModelList, err error)
	}
	IUser interface {
		List(ctx context.Context, inp *input.UserInput) (records []input.UserModelList, total int, err error)
		Detail(ctx context.Context, id int) (records input.UserModelDetail, err error)
		Edit(ctx context.Context, inp *input.UserEditInput) (err error)
		Reset(ctx context.Context, username string) (err error)
	}
)

var (
	localLogin    ILogin
	localPlatForm IPlatForm
	localUser     IUser
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

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
