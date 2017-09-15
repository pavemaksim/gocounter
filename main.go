package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	toml "github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	Connection	string
	DB 			string
}

type Event struct {
	DocID		int
	Total		int
}

func main() {

	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()

	// increment counter
	router.GET("/counter/:id", func(c *gin.Context) {
		id := c.Param("id")
		session, err := mgo.Dial(config.Connection)
		if err != nil {
			panic(err)
		}

		defer session.Close()

		session.SetMode(mgo.Monotonic, true)

		col := session.DB(config.DB).C("Events")

		col.Upsert(
			bson.M{"docid": id},
			bson.M{"$inc": bson.M{"total": 1}},
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("ok"),
		})
	})

	router.Run(":8000")
}
