package article

import (
	dto "article/src/app/dto/article"
	models "article/src/infra/models/article"
	"context"
	"encoding/json"
	"log"
	"time"

	redis "article/src/infra/persistence/redis/article"

	"github.com/jmoiron/sqlx"
)

type ArticleRepository interface {
	Create(data *dto.ArticleReqDTO) error
	GetList(data *dto.GetArticleReqDTO) ([]*models.GetArticleModel, error)
}

const (
	CreateArticle     = `INSERT INTO article (author, title, body, created_at) values ($1, $2, $3, now())`
	GetArticle        = `SELECT id, author, title, body, created_at from article where deleted_at is null order by created_at desc`
	GetArticleByQuery = `SELECT id, author, title, body, created_at
	from article where id IN (select id from article where deleted_at is null)
	AND id IN (
	select id where title ILIKE '%' || $1 || '%'
	OR body ILIKE '%' || $1 || '%'
	)
	AND id IN (
		select id where author ILIKE '%' || $2 || '%'
	)
	order by created_at desc`
)

var statement PreparedStatement

type PreparedStatement struct {
	createArticle     *sqlx.Stmt
	getArticle        *sqlx.Stmt
	getArticleByQuery *sqlx.Stmt
}

type ArticleRepo struct {
	Connection *sqlx.DB
	Redis      redis.ArticleRedisInt
}

func NewArticleRepository(db *sqlx.DB, rds redis.ArticleRedisInt) ArticleRepository {
	repo := &ArticleRepo{
		Connection: db,
		Redis:      rds,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *ArticleRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *ArticleRepo) {
	statement = PreparedStatement{
		createArticle:     m.Preparex(CreateArticle),
		getArticle:        m.Preparex(GetArticle),
		getArticleByQuery: m.Preparex(GetArticleByQuery),
	}
}

func (p *ArticleRepo) Create(data *dto.ArticleReqDTO) error {

	_, err := statement.createArticle.Exec(data.Author, data.Title, data.Body)

	if err != nil {
		log.Println("Failed Query Create Article : ", err.Error())
		return err
	}

	return nil
}

func (p *ArticleRepo) GetList(data *dto.GetArticleReqDTO) ([]*models.GetArticleModel, error) {
	var resultData []*models.GetArticleModel
	var err error
	var key string

	if data.Author != "" || data.Query != "" {
		dataKey, _ := json.Marshal(data)
		key = string(dataKey)
	} else {
		key = "getAll"
	}

	dataRedis, err := p.Redis.GetData(context.Background(), key)
	if err != nil {
		log.Printf("unable to GET data from redis. error: %v", err)
	}

	if dataRedis != "" {
		// get data from redis if is there
		err = json.Unmarshal([]byte(dataRedis), &resultData)
		if err != nil {
			return nil, err
		}

		log.Println("data from redis")

	} else {
		if data.Author != "" || data.Query != "" {
			err = statement.getArticleByQuery.Select(&resultData, data.Query, data.Author)
		} else {
			err = statement.getArticle.Select(&resultData)
		}

		if err != nil {
			return nil, err
		}

		redisData, _ := json.Marshal(resultData)
		ttl := time.Duration(2) * time.Minute

		// set data to redis
		rds := p.Redis.SetData(context.Background(), key, redisData, ttl)
		if err := rds.Err(); err != nil {
			log.Printf("unable to SET data. error: %v", err)
		}
	}

	return resultData, nil
}
