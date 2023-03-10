package mock_dto

import (
	dto "article/src/app/dto/article"

	"github.com/stretchr/testify/mock"
)

type MockArticleDTO struct {
	mock.Mock
}

func NewMockArticleDTO() *MockArticleDTO {
	return &MockArticleDTO{}
}

var _ dto.ArticleDTOInterface = &MockArticleDTO{}

func (m *MockArticleDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
