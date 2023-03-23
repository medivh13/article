package article

import (
	mockDTO "article/mock/app/dto/article"
	mockRepo "article/mock/infra/repository/article"
	mockReponse "article/mock/interface/response"
	dto "article/src/app/dto/article"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

type MockArticleUseCase struct {
	mock.Mock
}

type ArticleUseCaseList struct {
	suite.Suite
	resp     *mockReponse.MockResponse
	mockDTO  *mockDTO.MockArticleDTO
	mockRepo *mockRepo.MockArticleRepo
	useCase  ArticleUCInterface

	dtoCreate     *dto.ArticleReqDTO
	dtoCreateFail *dto.ArticleReqDTO
	dtoGet        *dto.GetArticleReqDTO
}

func (suite *ArticleUseCaseList) SetupTest() {
	suite.resp = new(mockReponse.MockResponse)
	suite.mockDTO = new(mockDTO.MockArticleDTO)
	suite.mockRepo = new(mockRepo.MockArticleRepo)
	suite.useCase = NewArticleUseCase(suite.mockRepo)

	suite.dtoCreate = &dto.ArticleReqDTO{
		Author: "test",
		Title:  "test",
		Body:   "test",
	}

	suite.dtoCreateFail = &dto.ArticleReqDTO{
		Author: "test",
		Title:  "test",
	}

	suite.dtoGet = &dto.GetArticleReqDTO{
		Query:  "article",
		Author: "john",
	}
}

func (u *ArticleUseCaseList) TestCreateSuccess() {
	u.mockRepo.Mock.On("Create", mock.Anything).Return(nil)
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(nil, err)
}

func (u *ArticleUseCaseList) TestCreateFail() {
	u.mockRepo.Mock.On("Create", mock.Anything).Return(errors.New(mock.Anything))
	err := u.useCase.Create(u.dtoCreate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *ArticleUseCaseList) TestGetListSuccess() {
	u.mockRepo.Mock.On("GetList", mock.Anything).Return(mock.Anything, nil)
	_, err := u.useCase.GetList(u.dtoGet)
	u.Equal(nil, err)
}

func (u *ArticleUseCaseList) TestGetListFail() {
	u.mockRepo.Mock.On("GetList", mock.Anything).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetList(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(ArticleUseCaseList))
}
