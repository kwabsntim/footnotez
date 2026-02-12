package internal

import (
	"fmt"
	"strings"
)

type ArticleService struct {
	repo Repository
}

//the below is the function that generates the slug

func generateSlug(title string) string {
	//converting the title to lowercase
	slug := strings.ToLower(title)
	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove special characters (keep only alphanumeric and hyphens)
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}
	slug = result.String()
	// Replace multiple hyphens with a single one
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	return slug
}

func NewArticleService(repo Repository) ArticleInterface {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) CreateArticle(article *Article) error {
	if article.Title == "" || article.Content == "" {
		return fmt.Errorf("title and content cannot be empty")

	}
	article.Slug = generateSlug(article.Title)
	err := s.repo.Create(article)
	if err != nil {
		return fmt.Errorf("failed to create article: %v", err)
	}
	return nil

}
func (s *ArticleService) GetArticle(id int) (*Article, error) {
	return s.repo.GetByID(id)
}

func (s *ArticleService) GetAllArticles() ([]Article, error) {
	return s.repo.GetAll()
}

func (s *ArticleService) UpdateArticle(id int, title, content string) (*Article, error) {
	article, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		article.Title = title
		article.Slug = generateSlug(title)
	}
	if content != "" {
		article.Content = content
	}

	err = s.repo.Update(article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.repo.Delete(id)
}
