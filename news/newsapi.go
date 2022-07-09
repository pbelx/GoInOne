ackage news

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	qlog "serverlessRadio/database"

	"github.com/gin-gonic/gin"
)

type JsonReply struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"media,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Language      string `xml:"language"`
		WebMaster     string `xml:"webMaster"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Description   string `xml:"description"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
			Source      struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"source"`
		} `xml:"item"`
	} `xml:"channel"`
}

//news Api
func NewsApi(c *gin.Context) {

	qlog.Getip(c)
	url, _ := os.ReadFile("./news/url.txt")
	resp, err := http.Get(string(url))
	if err != nil {
		fmt.Println("got Err")
		c.JSON(200, gin.H{
			"message": "unable to get Data",
		})
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err")
	}
	var news Rss
	var snews []JsonReply

	xml.Unmarshal(data, &news)
	for i, v := range news.Channel.Item {
		fmt.Println(v.Title)
		fmt.Println(v.Link)
		fmt.Println(i)
		g := JsonReply{v.Title, v.Link}
		snews = append(snews, g)

		fmt.Println("*******")

	}
	fmt.Println(len(snews))
	c.JSON(http.StatusOK, gin.H{
		"znews":  snews,
		"length": len(snews),
	})

}
