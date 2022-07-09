package radio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	qlog "serverlessRadio/database"

	"github.com/gin-gonic/gin"
)

// read radio stations

type Rstations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Xstations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func RadioReply(url string) gin.HandlerFunc {

	funcRadio := func(c *gin.Context) {
		qlog.Getip(c)
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "service offline",
			})
		} else {
			io.Copy(c.Writer, resp.Body)
		}
	}
	return gin.HandlerFunc(funcRadio)

}

func KenyaRedirect(url string) gin.HandlerFunc {

	funcRadio := func(c *gin.Context) {
		// getip(c)
		qlog.Getip(c)
		c.Redirect(302, url)
	}
	return gin.HandlerFunc(funcRadio)

}
func KenyaListen(c *gin.Engine) {

	fileSpecial, _ := os.ReadFile("./radiomodule/redirectStations.json")
	var kstation []Xstations

	json.Unmarshal([]byte(fileSpecial), &kstation)

	for _, k := range kstation {
		var kurl = k.Name
		c.GET("/"+kurl, KenyaRedirect(k.Url))
	}
}

func LocalradioListen(c *gin.Engine) {

	filex, err := os.ReadFile("./radiomodule/stations.json")
	if err != nil {
		fmt.Println(err)
	}
	var xstation []Rstations

	json.Unmarshal([]byte(filex), &xstation)

	for _, v := range xstation {
		var zurl = v.Name

		c.GET("/"+zurl, RadioReply(v.Url))

	}

}
