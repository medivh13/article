package article

/*
 * Author      : Jody (github.com/medivh13)
 * Modifier    :
 * Domain      : article
 */
import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ArticleDTOInterface interface {
	Validate() error
}

type ArticleReqDTO struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (dto *ArticleReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Author, validation.Required),
		validation.Field(&dto.Title, validation.Required),
		validation.Field(&dto.Body, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetArticleReqDTO struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}

type GetArticleRespDTO struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
