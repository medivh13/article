package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

import (
	dto "article/src/app/dto/article"
	"context"
	"encoding/json"
	"log"
	"time"

	articleRepo "article/src/infra/persistence/postgres/article"

	redis "article/src/infra/persistence/redis/article"
)

type ArticleUCInterface interface {
	Create(data *dto.ArticleReqDTO) error
	GetList(req *dto.GetArticleReqDTO) ([]*dto.GetArticleRespDTO, error)
}

type articleUseCase struct {
	ArticleRepo articleRepo.ArticleRepository
	Redis       redis.ArticleRedisInt
}

func NewArticleUseCase(articleRepo articleRepo.ArticleRepository, rds redis.ArticleRedisInt) *articleUseCase {
	return &articleUseCase{
		ArticleRepo: articleRepo,
		Redis:       rds,
	}
}

func (uc *articleUseCase) Create(data *dto.ArticleReqDTO) error {

	err := uc.ArticleRepo.Create(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *articleUseCase) GetList(req *dto.GetArticleReqDTO) ([]*dto.GetArticleRespDTO, error) {
	var key string
	var resp []*dto.GetArticleRespDTO
	if req.Author != "" || req.Query != "" {
		dataKey, _ := json.Marshal(req)
		key = string(dataKey)
	} else {
		key = "getAll"
	}

	dataRedis, err := uc.Redis.GetData(context.Background(), key)
	if err != nil {
		log.Printf("unable to GET data from redis. error: %v", err)
	}

	if dataRedis != "" {
		// get data from redis if is there
		err = json.Unmarshal([]byte(dataRedis), &resp)
		if err != nil {
			return nil, err
		}

		log.Println("data from redis")

	} else {
		data, err := uc.ArticleRepo.GetList(req)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		resp = dto.ToArticle(data)

		redisData, _ := json.Marshal(resp)
		ttl := time.Duration(2) * time.Minute

		// set data to redis
		rds := uc.Redis.SetData(context.Background(), key, redisData, ttl)
		if err := rds.Err(); err != nil {
			log.Printf("unable to SET data. error: %v", err)

		}
	}

	return resp, nil
}
