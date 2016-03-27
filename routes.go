package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	u := c.Query("url")

	url, err := url.Parse(u)
	if err != nil {
		c.Status(400)
	}

	var book Book
	db.Where("host LIKE ? AND path LIKE ?", "%"+url.Host, url.Path+"%").first(&book)

	if book.Id != 0 {
		c.JSON(200, book)
	} else {
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
