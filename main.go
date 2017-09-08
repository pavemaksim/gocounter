package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	toml "github.com/BurntSushi/toml"
)

type Config struct {
	DBDsn string
}

func main() {

	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("mysql", config.DBDsn)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	router := gin.Default()

	// increment counter
	router.GET("/counter/:id", func(c *gin.Context) {
		id := c.Param("id")
		stmt, err := db.Prepare("insert into counters (id) values(?) on duplicate key update amount = amount + 1;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)

		if err != nil {
			fmt.Print(err.Error())
		}

		defer stmt.Close()

		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("ok"),
		})
	})

	router.Run(":8000")
}
