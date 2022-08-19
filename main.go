package main

import (
	"gin_todo/model"

	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	model.InitDB()
}

func main() {
	// Create instance
	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob("views/*.html")

	// routing
	r.GET("/", Index)
	r.GET("/:id", Show)
	r.GET("/new", New)
	r.POST("/create", Create)
	r.GET("/:id/edit", Edit)
	r.POST("/:id/update", Update)
	r.POST("/:id/delete", Delete)

	r.Run()
}

// controller
// index
func Index(ctx *gin.Context) {
	todos := model.GetAll()
	ctx.HTML(200, "index.html", gin.H{"todos": todos})
}

// show
func Show(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Redirect(302, "/")
	}
	todo := model.GetOne(id)
	ctx.HTML(200, "show.html", gin.H{"todo": todo})
}

// new
func New(ctx *gin.Context) {
	ctx.HTML(200, "new.html", gin.H{})
}

// create
func Create(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	model.CreateTodo(title, description)
	ctx.Redirect(302, "/")
}

// edit
func Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Redirect(302, "/")
	}
	todo := model.GetOne(id)
	ctx.HTML(200, "edit.html", gin.H{"todo": todo})
}

// update
func Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	model.UpdateTodo(id, title, description)
	ctx.Redirect(302, "/")
}

// delete
func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	model.DeleteTodo(id)
	ctx.Redirect(302, "/")
}
