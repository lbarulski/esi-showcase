package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.GET("/:seconds", func(c *gin.Context) {
		delay, _ := strconv.Atoi(c.Param("seconds"))
		time.Sleep(time.Duration(delay) * time.Second)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.GET("/delay/:seconds/:color", func(c *gin.Context) {
		delay, _ := strconv.Atoi(c.Param("seconds"))
		time.Sleep(time.Duration(delay) * time.Second)

		c.Header("Cache-Control", "no-cache, no-store")

		c.HTML(http.StatusOK, "delay.tmpl", gin.H{
			"delay": delay,
			"color": c.Param("color"),
		})
	})
	r.Run()
}
