package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderHtml(html string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, html, nil)
	}
}
