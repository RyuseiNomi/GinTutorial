package main

import (
	"github.com/RyuseiNomi/GinTutorial/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// DBの初期化
	db.Init()

	// Index
	router.GET("/", func(ctx *gin.Context) {
		todos := db.GetAll()
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	// Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		db.Insert(text, status)
		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := db.FetchOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	// Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		db.Update(id, text, status)
		ctx.Redirect(302, "/")
	})

	// Confirm Delete
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := db.FetchOne()
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	// Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		db.Delete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
