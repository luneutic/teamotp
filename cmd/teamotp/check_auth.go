package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) bool {
	session := sessions.Default(c)
	return session.Get("username") != nil
}

func Unauthorized_404(c *gin.Context) bool {
	if !IsAuthenticated(c) {
		c.AbortWithStatus(http.StatusNotFound)
		return true
	}

	return false
}

func Default_404() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		code := c.Writer.Status()

		if code == http.StatusNotFound {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
		}
	}
}
