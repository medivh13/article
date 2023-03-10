package article

import "time"

type GetArticleModel struct {
	ID        int64     `db:"id"`
	Author    string    `db:"author"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}
