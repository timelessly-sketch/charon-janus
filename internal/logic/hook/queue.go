package hook

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

type queueImpl struct {
	ch chan entity.SysLog
}

const (
	maxBatchSize = 10
	maxWaitTime  = 1 * time.Second
)

var (
	queue *queueImpl
)

func init() {
	queue = &queueImpl{
		ch: make(chan entity.SysLog, 100),
	}
	go queue.consumer(gctx.New())
}

func Queue() *queueImpl {
	return queue
}

// Push 推送日志到队列
func (q *queueImpl) Push(ctx context.Context, log entity.SysLog) {
	select {
	case q.ch <- log: // 非阻塞写入
	default:
		g.Log().Warningf(ctx, "Log queue is full, discarding logs: %s", log.ReqId)
	}
}

// consumer 消费者协程（批量处理）
func (q *queueImpl) consumer(ctx context.Context) {
	var (
		logs  []entity.SysLog
		timer = time.NewTimer(maxWaitTime)
	)
	defer timer.Stop()

	for {
		select {
		case log := <-q.ch:
			logs = append(logs, log)
			if len(logs) >= maxBatchSize {
				q.flush(ctx, logs)
				logs = nil
				timer.Reset(maxWaitTime)
			}

		case <-timer.C:
			if len(logs) > 0 {
				q.flush(ctx, logs)
				logs = nil
			}
			timer.Reset(maxWaitTime)
		}
	}
}

// flush 批量写入数据库
func (q *queueImpl) flush(ctx context.Context, logs []entity.SysLog) {
	if _, err := dao.SysLog.Ctx(ctx).Data(logs).Insert(); err != nil {
		g.Log().Errorf(ctx, "batch insertion of logs failed: %v, this time affecting %d entries", err, len(logs))
		return
	}
	g.Log().Debugf(ctx, "successful batch insertion of %d logs", len(logs))
	return
}
