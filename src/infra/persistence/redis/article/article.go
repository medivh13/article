package article

import (
	"time"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type ArticleRedis struct {
	Rdb *redis.Client
}

type ArticleRedisInt interface {
	SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
	GetData(ctx context.Context, key string) (string, error)
}

func NewArticleRedis(rdb *redis.Client) *ArticleRedis {
	return &ArticleRedis{
		Rdb: rdb,
	}
}

func (p *ArticleRedis) SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd {
	rds := p.Rdb.Set(context.Background(), key, value, ttl)

	return rds
}

func (p *ArticleRedis) GetData(ctx context.Context, key string) (string, error) {
	dataRedis, err := p.Rdb.Get(context.Background(), key).Result()

	return dataRedis, err
}
