package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	c "github.com/irnak4t/leaderboards/controllers"
	"github.com/irnak4t/leaderboards/middleware"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"inc":         middleware.Inc,
		"parseRecord": middleware.ParseRecord,
	})
	router.LoadHTMLGlob("templates/*.html")

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
	router.Run()
}
