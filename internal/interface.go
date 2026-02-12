package internal

type Articles interface {
	Create(article *Article) error
	GetByID(id int) (*Article, error)
	Update(article *Article) error
	Delete(id int) error
	GetAll() ([]*Article, error)
}

type ArticleService struct {
}
