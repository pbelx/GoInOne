package caa

import (
	"os"

	qlog "serverlessRadio/database"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func FlightF(c *gin.Context) {
	// getip(c)
	qlog.Getip(c)

	type Flights struct {
		Airline    string `json:"airline,omitempty"`
		FlightCode string `json:"flightCode,omitempty"`
		ET         string `json:"expectedtime,omitempty"`
		ORIGIN     string `json:"origin,omitempty"`
		Status     string `json:"status,omitempty"`
		Category   string `json:"category,omitempty"`
	}
	type Flights2 struct {
		Airline     string `json:"airline,omitempty"`
		FlightCode  string `json:"flightCode,omitempty"`
		ET          string `json:"expectedtime,omitempty"`
		Destination string `json:"destination,omitempty"`
		Status      string `json:"status,omitempty"`
		Category    string `json:"category,omitempty"`
	}

	Flightlist := make([]Flights, 0)
	DFlightlist := make([]Flights2, 0)
	cc := colly.NewCollector()
	cc.OnHTML(".fids_table", func(e *colly.HTMLElement) {
		item := Flights{}
		itemx := Flights2{}
		// item.Category = e.ChildText("th:nth-child(4)")
		category := e.ChildText("th:nth-child(5)")
		// airline := e.ChildText("td:nth-child(1)")
		// fcode := e.ChildText("td:nth-child(2)")
		// status := e.ChildText("td:nth-child(6)")

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			if category == "ETA" {
				if el.ChildText("td:nth-child(1)") != "" {
					item.Category = el.ChildText("th:nth-child(4)")
					item.Airline = el.ChildText("td:nth-child(1)")
					// fmt.Println(item.Airline)
					item.Status = el.ChildText("td:nth-child(6)")
					item.ORIGIN = el.ChildText("td:nth-child(4)")
					item.FlightCode = el.ChildText("td:nth-child(3)")
					item.ET = el.ChildText("td:nth-child(5)")
					Flightlist = append(Flightlist, item)
					// item.Category = category
					// item.Airline = airline
					// item.Status = status
					// item.FlightCode = fcode
					// Flightlist = append(Flightlist, item)
				}
			} else if category == "ETD" {
				if el.ChildText("td:nth-child(1)") != "" {
					itemx.Category = el.ChildText("th:nth-child(4)")
					itemx.Airline = el.ChildText("td:nth-child(1)")

					itemx.Status = el.ChildText("td:nth-child(6)")
					itemx.Destination = el.ChildText("td:nth-child(4)")
					itemx.FlightCode = el.ChildText("td:nth-child(3)")
					itemx.ET = el.ChildText("td:nth-child(5)")
					DFlightlist = append(DFlightlist, itemx)
				}
			}
		})
	})
	url, _ := os.ReadFile("./caa/caa.txt")
	cc.Visit(string(url))

	c.JSON(200, gin.H{
		"data":  Flightlist,
		"data2": DFlightlist,
	})
}
