package main

import (
	"github.com/gin-gonic/gin"
	"go-blog-example/_day03/controllers"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	app := App{
		WebEngine:         gin.Default(),
		ArticleController: controllers.NewArticle(),
	}
	_ = app.Run()
}
