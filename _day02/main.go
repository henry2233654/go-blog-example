package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	DB = InitialDb()
	webEngine := InitialWebEngine()
	err := webEngine.Run(":8081")
	if err != nil {
		panic(err)
	}
}

func InitialDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./temp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&Article{})
	return db
}

func InitialWebEngine() *gin.Engine {
	e := gin.Default()
	articles := e.Group("/articles")
	articles.GET("", ListArticles)
	articles.POST("", CreateArticle)
	articles.GET("/:id", GetArticle)
	articles.PUT("/:id", UpdateArticle)
	articles.DELETE("/:id", DeleteArticle)
	return e
}

type Article struct {
	ID        uint           `json:"id" gorm:"column:id;primaryKey"`
	Title     string         `json:"title" gorm:"column:title"`
	Content   string         `json:"content" gorm:"column:content"`
	Author    string         `json:"author" gorm:"column:author"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func ListArticles(c *gin.Context) {
	var articles []*Article
	err := DB.Find(&articles).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	id, err := GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	var a Article
	err = DB.First(&a, id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
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
	a.ID = 0
	err = DB.Save(&a).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func UpdateArticle(c *gin.Context) {
	id, err := GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	var a Article
	err = DB.First(&a, id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	a.ID = id
	err = DB.Save(&a).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func DeleteArticle(c *gin.Context) {
	id, err := GetIdFromUri(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be uint.")
		return
	}
	var a Article
	err = DB.First(&a, id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusNotFound, "Can't find the article.")
		return
	}
	err = DB.Delete(&a).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func GetIdFromUri(c *gin.Context) (id uint, err error) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	id = uint(idP)
	return
}
