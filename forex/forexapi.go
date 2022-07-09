package forex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	qlog "serverlessRadio/database"
)

func Forexrates(c *gin.Context) {
	qlog.Getip(c)
	type rates struct {
		EUR float64 `json:"EUR"`
		GBP float64 `json:"GBP"`
		KES float64 `json:"KES"`
		UGX float64 `json:"UGX"`
	}

	type jsondata struct {
		Base       string `json:"base"`
		Disclaimer string `json:"disclaimer"`
		License    string `json:"license"`
		Rates      rates  `json:"rates"`
		Timestamp  string `json:"timestamp"`
	}

	fmt.Println("Fetching Rates now...")
	url, _ := os.ReadFile("./forex/url.txt")
	// urlR,_ = ioutil.ReadAll(url)
	resp, err := http.Get(string(url))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		reply, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "service offline 404",
			})
		}

		// fmt.Println(string(reply))

		var xrates jsondata
		json.Unmarshal([]byte(reply), &xrates)
		//get ugx rate
		realugx := math.Floor(float64(xrates.Rates.UGX))
		//get other currencies compared to ugx
		// fmt.Println(realugx)

		Euros := math.Floor(realugx / xrates.Rates.EUR)
		Pounds := math.Floor(realugx / xrates.Rates.GBP)
		Ksh := math.Floor(realugx / xrates.Rates.KES)

		// fmt.Printf("Euros %v ,Pounds %v,Ksh %v", Euros, Pounds, Ksh)
		c.JSON(200, gin.H{
			"Euros":  Euros,
			"Pounds": Pounds,
			"KSH":    Ksh,
			"USD":    realugx,
		})
	}
}
