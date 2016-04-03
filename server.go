package main

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	v1 := r.Group("api")

	v1.GET("/book", GetBook)
	r.Run(":8080")
}

func scrape(u string, ch chan<- Book) {
	url, err := url.Parse(u)
	doc, err := goquery.NewDocument(u)
	if err != nil {
		log.Fatal(err)
	}
	img := doc.Find("#imgBlkFront").AttrOr("data-a-dynamic-image", "not found")
	var author []string
	doc.Find("#byline").Find(".author .contributorNameID").Each(func(i int, s *goquery.Selection) {
		author = append(author, s.Text())
	})
	r, _ := regexp.Compile("\"(.*?)\"")
	book := Book{
		Title:   doc.Find("#productTitle").Text(),
		Author:  strings.Join(author, ", "),
		Img:     r.FindString(img),
		FullUrl: u,
		Host:    url.Host,
		Path:    url.Path,
	}
	db.Create(&book)
	ch <- book
}

func GetBook(c *gin.Context) {
	u := c.Query("url")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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

	log.Println("Received", url.Path)

	if book.ID != 0 {
		log.Print("Found in DB ", book.ID)
		c.JSON(200, book)
		return
	} else {
		ch := make(chan Book)
		go scrape(u, ch)
		c.JSON(200, <-ch)
	}
}
