package internal

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Articles {
	return &Repository{DB: db}
}

// creating the crud for the article
func (r *Repository) Create(article *Article) error {
	query := `INSERT INTO articles (title,content,slug) VALUES(?,?,?)`
	result, err := r.DB.Exec(query, article.Title, article.Content, article.Slug)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	article.ID = int(id)
	return nil

}
func (r *Repository) GetAll() ([]*Article, error) {
	query := `SELECT id, title, content, slug, created_at FROM articles ORDER BY created_at DESC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article
	for rows.Next() {
		article := &Article{}
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.Slug,
			&article.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
func (r *Repository) GetByID(id int) (*Article, error) {
	query := `SELECT id,title,content,slug,created_at FROM articles WHERE id=?`
	article := &Article{}
	err := r.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.Slug, &article.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("article not found")
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (r *Repository) Update(article *Article) error {
	query := `UPDATE articles SET title = ?, content = ?, slug = ? WHERE id = ?`
	_, err := r.DB.Exec(query, article.Title, article.Content, article.Slug, article.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM articles WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
