package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

import (
	"encoding/json"
	"net/http"

	dto "article/src/app/dto/article"
	usecases "article/src/app/usecases/article"
	common_error "article/src/infra/errors"
	"article/src/interface/rest/response"
)

type ArticleHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
}

type articleHandler struct {
	response response.IResponseClient
	usecase  usecases.ArticleUCInterface
}

func NewArticleHandler(r response.IResponseClient, h usecases.ArticleUCInterface) ArticleHandlerInterface {
	return &articleHandler{
		response: r,
		usecase:  h,
	}
}

func (h *articleHandler) Create(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.ArticleReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Create(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Adding New Article",
		nil,
		nil,
	)
}

func (h *articleHandler) GetList(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetArticleReqDTO{}

	if r.URL.Query().Get("query") != "" {
		getDTO.Query = r.URL.Query().Get("query")
	}

	if r.URL.Query().Get("author") != "" {
		getDTO.Author = r.URL.Query().Get("author")
	}

	data, err := h.usecase.GetList(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Article",
		data,
		nil,
	)
}
