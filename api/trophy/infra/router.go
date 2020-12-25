package infra

import (
	"github.com/asatake/trophies-management/api/trophy/interfaces/controller"
	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func route() {
	router := gin.Default()

	contentController := controller.NewContentController(NewSqlHandler())

	router.GET("/:id/", func(c *gin.Context) { contentController.Show(c) })

	Router = router
}
