package models

import (
	"go-blog-example/_day03/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initialDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./temp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&Article{})
	return db
}

type Article struct {
	db *gorm.DB
}

func NewArticle() *Article {
	a := new(Article)
	a.db = initialDb()
	return a
}

func (model *Article) ListArticles() ([]*entities.Article, error) {
	var articles []*entities.Article
	err := model.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (model *Article) GetArticle(id uint) (*entities.Article, error) {
	var article entities.Article
	err := model.db.First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (model *Article) CreateArticle(article entities.Article) error {
	article.ID = 0
	return model.db.Save(&article).Error
}

func (model *Article) UpdateArticle(id uint, title, content, author string) error {
	article, err := model.GetArticle(id)
	if err != nil {
		return err
	}
	article.Title = title
	article.Content = content
	article.Author = author
	return model.db.Save(article).Error
}

func (model *Article) DeleteArticle(id uint) error {
	article, err := model.GetArticle(id)
	if err != nil {
		return err
	}
	return model.db.Delete(article).Error
}
