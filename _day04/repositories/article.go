package repositories

import (
	"go-blog-example/_day04/entities"
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

func (repo *Article) ListArticles() ([]*entities.Article, error) {
	var articles []*entities.Article
	err := repo.db.Find(&articles).Error
	return articles, err
}

func (repo *Article) GetArticle(id uint) (*entities.Article, error) {
	var article entities.Article
	err := repo.db.First(&article, id).Error
	return &article, err
}

func (repo *Article) SaveArticle(article *entities.Article) error {
	return repo.db.Save(&article).Error
}

func (repo *Article) DeleteArticle(article *entities.Article) error {
	return repo.db.Delete(article).Error
}
