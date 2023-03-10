package article

import (
	dto "article/src/app/dto/article"
	models "article/src/infra/models/article"
	"log"

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
}

func NewArticleRepository(db *sqlx.DB) ArticleRepository {
	repo := &ArticleRepo{db}
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
	if data.Author != "" || data.Query != "" {
		err = statement.getArticleByQuery.Select(&resultData, data.Query, data.Author)
	} else {
		err = statement.getArticle.Select(&resultData)
	}

	if err != nil {
		return nil, err
	}

	return resultData, nil
}
