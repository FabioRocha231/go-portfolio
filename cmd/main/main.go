package main

import (
	"github.com/FabioRocha231/go-portfolio/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("css", "../../static/css")
	r.LoadHTMLGlob("../../static/*.html")

	r.GET("/", controllers.RenderHtml("index.html"))

	r.Run()
}
