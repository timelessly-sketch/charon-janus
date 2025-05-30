package system

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type sLog struct{}

func NewLog() *sLog {
	return &sLog{}
}

func init() {
	service.RegisterLog(NewLog())
}

func (s *sLog) List(ctx context.Context, inp *input.LogInput) (records []input.LogRecords, total int, err error) {
	var (
		db   = dao.SysLog.Ctx(ctx)
		cols = dao.SysLog.Columns()
	)

	db = db.OmitEmpty().Where(g.Map{cols.ClientIp: inp.ClientIp}).WhereLike(cols.Username, "%"+inp.UserName+"%").WhereLike(cols.Url, "%"+inp.Path+"%")
	err = db.Page(inp.Page, inp.Size).ScanAndCount(&records, &total, true)
	return
}
