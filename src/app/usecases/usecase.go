package usecases

import (
	articleUC "article/src/app/usecases/article"
)

type AllUseCases struct {
	ArticleUC articleUC.ArticleUCInterface
}
