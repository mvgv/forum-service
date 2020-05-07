package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mvgv/forum-service/controller"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func urlMapping() {
	router.GET("/boards/:board_id/topics/:topic_id/posts", controller.GetPosts)
}

/*StartApplication inicializa as rotas do webserver*/
func StartApplication() {

	urlMapping()
	err := router.Run(":8081")

	if err != nil {
		panic(err)
	}
}
