package article

import models "article/src/infra/models/article"

func ToArticle(datas []*models.GetArticleModel) []*GetArticleRespDTO {
	var resp []*GetArticleRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnArticle(m))
	}
	return resp
}

func ToReturnArticle(d *models.GetArticleModel) *GetArticleRespDTO {
	return &GetArticleRespDTO{
		ID:        d.ID,
		Author:    d.Author,
		Title:     d.Title,
		Body:      d.Body,
		CreatedAt: d.CreatedAt,
	}
}
