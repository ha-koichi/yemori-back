package main



import (
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/PuerkitoBio/goquery"
)

func handler(c *gin.Context) {
	fmt.Fprint(c.Writer, "sometshing")
}

func scryping(c *gin.Context) {
	url := c.Query("url")
	doc, err := goquery.NewDocument(url)
	if err != nil {
			panic(err)
	}
	var doms []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		doms = append(doms, link)
	})
	jsons, err := json.Marshal(doms)
	if err != nil {
			return
	}
	fmt.Fprint(c.Writer, string(jsons))
}

func main() {
	router := gin.Default()

	router.GET("/hello", handler)
	router.GET("/q", scryping)

	router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
					"message": "ping!",
			})
	})
	router.Run()
}