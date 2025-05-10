package cache

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

var cache *gcache.Cache

func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

func SetAdapter(ctx context.Context) {
	var adapter gcache.Adapter

	switch g.Cfg().MustGet(ctx, "cache.adapter").String() {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	default:
		adapter = gcache.NewAdapterMemory()
	}

	g.DB().GetCache().SetAdapter(adapter)

	cache = gcache.New()
	cache.SetAdapter(adapter)
	g.Log().Info(ctx, "cache initialized")
}

func RemoveByPrefix(ctx context.Context, prefix string) (err error) {
	keys, err := cache.Keys(ctx)
	if err != nil {
		return
	}
	for _, v := range keys {
		key := gconv.String(v)
		if strings.HasPrefix(key, prefix) {
			_, _ = cache.Remove(ctx, key)
		}
	}
	return nil
}
