package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irnak4t/leaderboards/errors"
	"github.com/irnak4t/leaderboards/models"
	"github.com/irnak4t/leaderboards/parse"
	"github.com/irnak4t/leaderboards/sort"
)

type RecordController struct{}

var sorted map[string]map[string][]models.Record

func (rc RecordController) Index(ctx *gin.Context) {
	records := models.Record{}.GetAll()
	var isEmpty bool

	if len(records) == 0 {
		isEmpty = true
	} else {
		sort.SortByRecord(records, &sorted)
	}
	ctx.HTML(200, "index.html", gin.H{"data": sorted, "isEmpty": isEmpty})
}

func (rc RecordController) Show(ctx *gin.Context) {
	runner := ctx.Param("runner")
	records := models.Record{}.GetRuns(runner)

	if len(records) == 0 {
		ctx.AbortWithStatus(404)
	} else {
		sort.SortByRecord(records, &sorted)
		ctx.HTML(200, "show.html", gin.H{"data": sorted, "runner": runner})
	}
}

func (rc RecordController) Add(ctx *gin.Context) {
	ctx.HTML(200, "add.html", gin.H{})
}

func (rc RecordController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	errors.FailOnError(err)

	record := models.Record{}.Get(idInt)

	record.Delete()
	ctx.Redirect(303, "/record")
}

func (rc RecordController) Store(ctx *gin.Context) {
	record := models.Record{
		Title:    ctx.PostForm("title"),
		Category: ctx.PostForm("category"),
		Runner:   ctx.PostForm("runner"),
		Url:      ctx.PostForm("url"),
	}
	record.Time = parse.ParseTime(ctx.PostForm("time"))

	record.Create()
	ctx.Redirect(303, "/record")
}
