package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

import (
	dto "article/src/app/dto/article"
	"log"

	articleRepo "article/src/infra/persistence/postgres/article"
)

type ArticleUCInterface interface {
	Create(data *dto.ArticleReqDTO) error
	GetList(req *dto.GetArticleReqDTO) ([]*dto.GetArticleRespDTO, error)
}

type articleUseCase struct {
	ArticleRepo articleRepo.ArticleRepository
}

func NewArticleUseCase(articleRepo articleRepo.ArticleRepository) *articleUseCase {
	return &articleUseCase{
		ArticleRepo: articleRepo,
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
	
	data, err := uc.ArticleRepo.GetList(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	dto.ToArticle(data)

	return dto.ToArticle(data), nil
}
