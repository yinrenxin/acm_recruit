package routers

import (
	"ACMZX/controller"
	"ACMZX/util"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetupRouter() *gin.Engine {
	//initialization page

	{
		r = gin.Default()
		// to solve the cross domain
		r.Use(util.CrosHandler())
	}
	// found templates
	//r.LoadHTMLGlob("templates/*")
	//
	//// found static
	//r.Static("/static", "static")

	//r.GET("/", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "acmIndex.html", nil)
	//})

	// join us
	{
		//r.GET("/join", func(ctx *gin.Context) {
		//	ctx.HTML(http.StatusOK, "login.html", nil)
		//})
		// submit form
		// create data

		r.POST("/join", controller.CreateForm)

	}

	// extract table
	{
		r.POST("/extract", controller.QueryExtractTable)
		r.POST("/admin", controller.CreateExtractTable)
	}

	return r
}