package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Getip(c *gin.Context) {
	dbs, err := sql.Open("sqlite3", "./database/logs.db")

	if err != nil {
		fmt.Println(err)
	}
	defer dbs.Close()
	uaj := c.GetHeader("User-Agent")
	urlx := c.FullPath()
	ipx := c.ClientIP()
	timex := time.Now()
	fmt.Println(uaj, urlx, ipx, timex)
	stmt := "INSERT INTO zlog (useragent,ip,url,time) VALUES (?,?,?,?)"
	dbs.Prepare(stmt)
	fmt.Println(ipx)

	dbs.Exec(stmt, uaj, ipx, urlx, timex)

	rows, err := dbs.Query("SELECT useragent,ip FROM zlog")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {

		var useragent string
		var ip string
		rows.Scan(&useragent, &ip)
		// fmt.Println("w00t")
		// fmt.Printf(("%v,%v\n"), useragent, ip)

	}

}
