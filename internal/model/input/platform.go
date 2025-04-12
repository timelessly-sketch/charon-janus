package input

import "charon-janus/internal/model/entity"

type PlatFormInput struct {
	PlatformName string `json:"platform_name"`
	PlatformCode string `json:"platform_code"`
	PageReq
}

type PlatFormModelList struct {
	entity.SysPlatform
}

type PlatFormEditInput struct {
	entity.SysPlatform
}
