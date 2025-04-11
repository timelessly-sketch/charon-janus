package system

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
)

type sPlatForm struct{}

func NewPlatForm() *sPlatForm {
	return &sPlatForm{}
}

func init() {
	service.RegisterPlatForm(NewPlatForm())
}

func (s *sPlatForm) List(ctx context.Context, inp *input.PlatFormInput) (records []input.PlatFormModelList, total int, err error) {
	cols := dao.SysPlatform.Columns()
	db := dao.SysPlatform.Ctx(ctx)
	if inp.PlatformCode != "" {
		db = db.WhereLike(cols.PlatformCode, "%"+inp.PlatformCode+"%")
	}
	if inp.PlatformName != "" {
		db = db.Where(cols.PlatformName, "%"+inp.PlatformName+"%")
	}
	err = db.Page(inp.Page, inp.Size).ScanAndCount(&records, &total, true)
	return
}
