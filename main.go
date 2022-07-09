package main

import (
	"fmt"

	Fl "serverlessRadio/caa"         //flights
	Fx "serverlessRadio/forex"       //forex
	News "serverlessRadio/news"      //news
	Qq "serverlessRadio/quotes"      //quotes
	Rl "serverlessRadio/radiolist"   //radiolist
	Rm "serverlessRadio/radiomodule" //Radio Module

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	fmt.Println("serverLess up")

	r := gin.Default()
	r.Use(cors.Default())

	Rm.LocalradioListen(r) //Local radios listen
	Rm.KenyaListen(r)      //Redirect Stations
	r.GET("/", Rl.Rlist)
	r.GET("/quotes", Qq.Quotes)
	r.GET("/forex", Fx.Forexrates)
	r.GET("/flights", Fl.FlightF)
	r.GET("/news", News.NewsApi)
	// r.GET("/time", Lg.Getip)

	// r.StaticFS("/dww", http.Dir("/var/www/html/tech/dw/"))

	// go func() {
	// 	r.Run(":8080")
	// }()
	// r.RunTLS(":2053", "cert.pem", "key.pem")
	r.Run(":9000")
}
