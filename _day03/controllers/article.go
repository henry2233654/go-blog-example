package controllers

import (
	"github.com/gin-gonic/gin"
	"go-blog-example/_day03/entities"
	"go-blog-example/_day03/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Article struct {
	model *models.Article
}

func NewArticle() *Article {
	a := new(Article)
	a.model = models.NewArticle()
	return a
}

func (ctr *Article) Route(e *gin.Engine) {
	articles := e.Group("/articles")
	articles.GET("", ctr.ListArticles)
	articles.POST("", ctr.CreateArticle)
	articles.GET("/:id", ctr.GetArticle)
	articles.PUT("/:id", ctr.UpdateArticle)
	articles.DELETE("/:id", ctr.DeleteArticle)
}

func (ctr *Article) ListArticles(c *gin.Context) {
	articles, err := ctr.model.ListArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (ctr *Article) GetArticle(c *gin.Context) {
	id, err := ctr.GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	article, err := ctr.model.GetArticle(id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	c.JSON(http.StatusOK, article)
}

func (ctr *Article) CreateArticle(c *gin.Context) {
	var article entities.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = ctr.model.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ctr *Article) UpdateArticle(c *gin.Context) {
	id, err := ctr.GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	var article entities.Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = ctr.model.UpdateArticle(id, article.Title, article.Content, article.Author)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ctr *Article) DeleteArticle(c *gin.Context) {
	id, err := ctr.GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	err = ctr.model.DeleteArticle(id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ctr *Article) GetIdFromUri(c *gin.Context) (id uint, err error) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	id = uint(idP)
	return
}
