package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	c "github.com/irnak4t/leaderboards/controllers"
	"github.com/irnak4t/leaderboards/db"
	"github.com/irnak4t/leaderboards/errors"
	"github.com/irnak4t/leaderboards/middleware"
)

var router *gin.Engine

func main() {
	fileCheck()
	flagCheck()
	ginInit()
	ginRouting()
	router.Run()
}

func fileCheck() {
	cfgdir, err := os.UserConfigDir()
	errors.FailOnError(err)
	if _, err := os.Stat(cfgdir + "/lb.db.toml"); os.IsNotExist(err) {
		data, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/irnak4t/leaderboards/db/lb.db.toml.example")
		errors.FailOnError(err)
		err = ioutil.WriteFile(cfgdir+"/lb.db.toml", data, 0644)
	}
}

func flagCheck() {
	flag.Parse()
	if flag.Arg(0) == "migrate" {
		db.Migrate()
		os.Exit(0)
	}
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
