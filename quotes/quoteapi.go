package quotes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	qlog "serverlessRadio/database"
)

func Quotes(c *gin.Context) {
	// getip(c)
	qlog.Getip(c)

	type xquotes []struct {
		Q string `json:"q"`
		H string `json:"h"`
		A string `json:"a"`
	}

	url, _ := ioutil.ReadFile("./quotes/url.txt")
	resp, err := http.Get(string(url))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		bodydata, err := ioutil.ReadAll(resp.Body)
		// fmt.Printf(("%T"), bodydata)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "no data response",
			})
		}
		var qq xquotes

		json.Unmarshal([]byte(bodydata), &qq)
		author := qq[0].A
		quote := qq[0].Q

		c.JSON(200, gin.H{
			"Author": author,
			"Quote":  quote,
		})
	}

}
