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
	entity.AuthRole
}

type RoleDetailList struct {
	MenuList []entity.AuthMenu `json:"menuList"`
	MenuIds  []int             `json:"menuIds"`
}
