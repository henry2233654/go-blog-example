package main

import (
	"github.com/gin-gonic/gin"
	"go-blog-example/_day04/controllers"
)

type App struct {
	WebEngine         *gin.Engine
	ArticleController *controllers.Article
}

func (app App) Run() error {
	app.ArticleController.Route(app.WebEngine)
	return app.WebEngine.Run(":8081")
}
