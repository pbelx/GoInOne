package radiolist

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Xstations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Rlist(c *gin.Context) {
	var rr []Xstations
	var r2 []Xstations

	file, _ := os.ReadFile("./radiolist/stations.json")
	json.Unmarshal(file, &rr)

	file2, _ := os.ReadFile("./radiolist/redirectStations.json")
	json.Unmarshal(file2, &r2)
	var list = []string{}
	var list2 = []string{}
	for _, v := range rr {
		list = append(list, v.Name)
		// print(v.SName)
	}
	for _, k := range r2 {
		list2 = append(list2, k.Name)
		// print(v.SName)
	}

	// fmt.Println(list)

	c.JSON(200, gin.H{
		"list":  list,
		"list2": list2,
	})
}
