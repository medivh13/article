package article

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockDTO "article/mock/app/dto/article"
	mockUseCase "article/mock/app/usecases/article"
	mockReponse "article/mock/interface/response"
	dto "article/src/app/dto/article"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

type MockArticleHandler struct {
	mock.Mock
}

type ArticleHandlerList struct {
	suite.Suite
	resp    *mockReponse.MockResponse
	mockDTO *mockDTO.MockArticleDTO

	mockUseCase *mockUseCase.MockArticleUseCase
	h           ArticleHandlerInterface
	w           *httptest.ResponseRecorder

	dtoCreate     *dto.ArticleReqDTO
	dtoCreateFail *dto.ArticleReqDTO
	dtoGet        *dto.GetArticleReqDTO
}

func (suite *ArticleHandlerList) SetupTest() {
	suite.resp = new(mockReponse.MockResponse)
	suite.mockDTO = new(mockDTO.MockArticleDTO)
	suite.mockUseCase = new(mockUseCase.MockArticleUseCase)
	suite.h = NewArticleHandler(suite.resp, suite.mockUseCase)

	suite.w = httptest.NewRecorder()

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

func (s *ArticleHandlerList) TestCreateSuccess() {

	bodyBytes, _ := json.Marshal(s.dtoCreate)

	s.mockUseCase.Mock.On("Create", s.dtoCreate).Return(nil)
	s.resp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(nil)
	s.resp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	r := httptest.NewRequest("POST", "/articles", bytes.NewBuffer(bodyBytes))
	http.HandlerFunc(s.h.Create).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)

}

func (s *ArticleHandlerList) TestCreateFailDecode() {
	s.resp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(nil)
	r := httptest.NewRequest("POST", "/articles", nil)
	http.HandlerFunc(s.h.Create).ServeHTTP(s.w, r)
	s.Equal(400, s.w.Result().StatusCode)
}

func (s *ArticleHandlerList) TestCreateFailValidation() {
	bodyBytes, _ := json.Marshal(s.dtoCreateFail)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.resp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(nil)
	r := httptest.NewRequest("POST", "/create", bytes.NewBuffer(bodyBytes))
	http.HandlerFunc(s.h.Create).ServeHTTP(s.w, r)
	s.Equal(400, s.w.Result().StatusCode)
}

func (s *ArticleHandlerList) TestCreateFail() {
	bodyBytes, _ := json.Marshal(s.dtoCreate)
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("Create", s.dtoCreate).Return(errors.New(mock.Anything))
	s.resp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)
	r := httptest.NewRequest("POST", "/articles", bytes.NewBuffer(bodyBytes))
	http.HandlerFunc(s.h.Create).ServeHTTP(s.w, r)
	s.Equal(500, s.w.Result().StatusCode)
}

func (s *ArticleHandlerList) TestGetSuccess() {

	s.mockUseCase.Mock.On("GetList", s.dtoGet).Return(nil)
	s.resp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(nil)
	s.resp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	r := httptest.NewRequest("GET", "/articles?query=article&author=john", nil)
	http.HandlerFunc(s.h.GetList).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)

}

func (s *ArticleHandlerList) TestGetFail() {

	s.mockUseCase.Mock.On("GetList", s.dtoGet).Return(errors.New(mock.Anything))
	s.resp.Mock.On("HttpError", mock.Anything, mock.Anything).Return(nil)
	s.resp.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	r := httptest.NewRequest("GET", "/articles?query=article&author=john", nil)
	http.HandlerFunc(s.h.GetList).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)

}
func TestController(t *testing.T) {
	suite.Run(t, new(ArticleHandlerList))
}
