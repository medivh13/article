package mock_article

import (
	dto "article/src/app/dto/article"
	usecase "article/src/app/usecases/article"

	"github.com/stretchr/testify/mock"
)

type MockArticleUseCase struct {
	mock.Mock
}

func NewMockArticleUseCase() *MockArticleUseCase {
	return &MockArticleUseCase{}
}

var _ usecase.ArticleUCInterface = &MockArticleUseCase{}

func (m *MockArticleUseCase) Create(data *dto.ArticleReqDTO) error {
	args := m.Called(data)
	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *MockArticleUseCase) GetList(req *dto.GetArticleReqDTO) ([]*dto.GetArticleRespDTO, error) {
	args := m.Called(req)
	var (
		err  error
		resp []*dto.GetArticleRespDTO
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	if n, ok := args.Get(0).([]*dto.GetArticleRespDTO); ok {
		resp = n
	}

	return resp, err
}
