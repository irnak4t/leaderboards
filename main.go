package main

import (
	"flag"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	c "github.com/irnak4t/leaderboards/controllers"
	"github.com/irnak4t/leaderboards/db"
	"github.com/irnak4t/leaderboards/middleware"
)

var router *gin.Engine

func main() {
	flagCheck()
	ginInit()
	ginRouting()
	router.Run()
}

func ginInit() {
	router = gin.Default()
	router.SetFuncMap(template.FuncMap{
		"inc":         middleware.Inc,
		"parseRecord": middleware.ParseRecord,
	})
	router.LoadHTMLGlob("templates/*.html")
}

func ginRouting() {
	rc := c.RecordController{}
	rg := router.Group("/record")
	{
		rg.GET("", rc.Index)
		rg.GET("/add", rc.Add)
		rg.GET("/show/:runner", rc.Show)
		rg.POST("/store", rc.Store)
		rg.POST("/delete/:id", rc.Delete)
		rg.DELETE("/delete/:id", rc.Delete)
	}
	router.GET("", rc.Index)
}

func flagCheck() {
	flag.Parse()
	if flag.Arg(0) == "migrate" {
		db.Migrate()
		os.Exit(0)
	}
}
