package m_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Ts struct {
	rds *redis.Client
	ctx context.Context
}

func New(rds *redis.Client, ctx context.Context) *Ts {
	return &Ts{}
}

func (t *Ts) GetListLtrim(keyName string, count int64) ([]string, error) {
	return GetListLtrim(t.rds, t.ctx, keyName, count)
}
