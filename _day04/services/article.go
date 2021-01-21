package services

import (
	"go-blog-example/_day04/entities"
	"go-blog-example/_day04/repositories"
)

type Article struct {
	repo *repositories.Article
}

func NewArticle() *Article {
	a := new(Article)
	a.repo = repositories.NewArticle()
	return a
}

func (srv *Article) ListArticles() ([]*entities.Article, error) {
	articles, err := srv.repo.ListArticles()
	return articles, err
}

func (srv *Article) GetArticle(id uint) (*entities.Article, error) {
	article, err := srv.repo.GetArticle(id)
	return article, err
}

func (srv *Article) CreateArticle(article *entities.Article) error {
	article.ID = 0
	return srv.repo.SaveArticle(article)
}

func (srv *Article) UpdateArticle(id uint, title, content, author string) error {
	article, err := srv.GetArticle(id)
	if err != nil {
		return err
	}
	article.Title = title
	article.Content = content
	article.Author = author
	return srv.repo.SaveArticle(article)
}

func (srv *Article) DeleteArticle(id uint) error {
	article, err := srv.GetArticle(id)
	if err != nil {
		return err
	}
	return srv.repo.DeleteArticle(article)
}
