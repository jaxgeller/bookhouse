package main

import (
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	gorm.Model
	Title   string `form:"title"`
	Author  string `form:"author"`
	Host    string `gorm:"index" form:"host"`
	Path    string `gorm:"type:varchar(2083);index" form:"path"`
	Img     string `gorm:"type:varchar(2083);index" form:"img"`
	FullUrl string `gorm:"type:varchar(2083);index" form:"fullUrl"`
}

var db *gorm.DB
var err error

func main() {
	pgConnection := "postgres://bookhouseuser:bookhousepass@dockerhost:5432/bookhousedb?sslmode=disable"
	db, err = gorm.Open("postgres", pgConnection)
	if err != nil {
		panic("failed to connect database")
	}

	db.DB().SetMaxIdleConns(100)
	err = db.DB().Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}

	db.AutoMigrate(&Book{})
	db.Create(&Book{Title: "Da Vinci Code", Author: "Dan Brown", FullUrl: "http://www.amazon.com/Da-Vinci-Code-Dan-Brown/dp/0307474275", Img: "http://ecx.images-amazon.com/images/I/41cXJLj3BkL._SX277_BO1,204,203,200_.jpg", Host: "www.amazon.com", Path: "/Da-Vinci-Code-Dan-Brown/dp/0307474275/ref=sr_1_2"})

	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("api")

	v1.GET("/book", GetBook)
	v1.POST("/book", SaveBook)
	r.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetBook(c *gin.Context) {
	u := c.Query("url")
	if u == "" {
		c.Status(400)
		return
	}

	url, err := url.Parse(u)
	if err != nil {
		c.Status(400)
		return
	}

	var book Book
	db.Where("host LIKE ? AND path LIKE ?", "%"+url.Host, url.Path+"%").First(&book)

	if book.ID != 0 {
		log.Print("Book with ID ", book.ID, " found")
		c.JSON(200, book)
		return
	} else {
		log.Print(url.Path, " not found with host ", url.Host)
		// proxy := httputil.NewSingleHostReverseProxy(url)
		// proxy.ServeHTTP(c.Writer, c.Request)

	}
}

func SaveBook(c *gin.Context) {
	var book Book
	c.Bind(&book)

	db.Create(&book)
}
