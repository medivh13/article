package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : article
 */

import (
	dto "article/src/app/dto/article"
	models "article/src/infra/models/article"
	repo "article/src/infra/persistence/postgres/article"

	"github.com/stretchr/testify/mock"
)

var _ repo.ArticleRepository = &MockArticleRepo{}

type MockArticleRepo struct {
	mock.Mock
}

func (o *MockArticleRepo) Create(data *dto.ArticleReqDTO) error {
	args := o.Called(data)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (o *MockArticleRepo) GetList(data *dto.GetArticleReqDTO) ([]*models.GetArticleModel, error) {
	args := o.Called(data)

	var (
		err    error
		result []*models.GetArticleModel
	)

	if n, ok := args.Get(0).([]*models.GetArticleModel); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return result, err
}
