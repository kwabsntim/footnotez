package internal

type RepositoryInterface interface {
	Create(article *Article) error
	GetByID(id int) (*Article, error)
	Update(article *Article) error
	Delete(id int) error
	GetAll() ([]Article, error)
}

type ServiceInterface interface {
	CreateArticle(article *Article) error
	GetArticle(id int) (*Article, error)
	GetAllArticles() ([]Article, error)
	UpdateArticle(id int, title, content string) (*Article, error)
	DeleteArticle(id int) error
}
