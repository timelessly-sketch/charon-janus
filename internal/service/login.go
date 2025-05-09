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
		UserRoutes(ctx context.Context, code string) (records input.UserRoutes, err error)
		LoginMenuCacheKey(code string, id int) (key string)
	}
)

var (
	localLogin ILogin
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
