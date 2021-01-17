package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	articles := e.Group("/articles")
	articles.GET("", ListArticles)
	articles.POST("", CreateArticle)
	articles.GET("/:id", GetArticle)
	articles.PUT("/:id", UpdateArticle)
	articles.DELETE("/:id", DeleteArticle)
	err := e.Run(":8081")
	if err != nil {
		panic(err)
	}
}

// Article
type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var Articles = map[uint]*Article{
	1: &Article{ID: 1, Title: "Title 1", Content: "this is content 1.", Author: "Henry"},
	2: &Article{ID: 2, Title: "Title 2", Content: "this is content 2.", Author: "Amy"},
	3: &Article{ID: 3, Title: "Title 3", Content: "this is content 3.", Author: "Chris"},
}

var NewestID = uint(3)

func ListArticles(c *gin.Context) {
	resp := make([]*Article, len(Articles))
	i := 0
	for _, a := range Articles {
		resp[i] = a
		i++
	}
	c.JSON(http.StatusOK, resp)
}

func GetArticle(c *gin.Context) {
	id, err := GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	a, ok := Articles[id]
	if !ok {
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	c.JSON(http.StatusOK, a)
}

func CreateArticle(c *gin.Context) {
	var a Article
	err := c.BindJSON(&a)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	NewestID++
	a.ID = NewestID
	Articles[NewestID] = &a
	c.JSON(http.StatusNoContent, nil)
}

func UpdateArticle(c *gin.Context) {
	id, err := GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	var a Article
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	a.ID = id
	Articles[id] = &a
	c.JSON(http.StatusNoContent, nil)
}

func DeleteArticle(c *gin.Context) {
	id, err := GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	_, ok := Articles[id]
	if !ok {
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	delete(Articles, id)
}

func GetIdFromUri(c *gin.Context) (id uint, err error) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	id = uint(idP)
	return
}
