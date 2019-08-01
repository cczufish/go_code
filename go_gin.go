package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
)

type Person struct {
	Id   int
	Name string
}


func FetchSingleUser(c *gin.Context) {

	id := c.Param("id")

	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	var (
		person Person
		result gin.H
	)
	row := db.QueryRow("select id, name from user_info where id = ?;", id)
	err = row.Scan(&person.Id, &person.Name)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/userinfo") // http://localhost:8080/api/v1/userinfo/:1
	{
		//v1.POST("/", CreateUser)
		//v1.GET("/", FetchAllUsers)
		v1.GET("/:id", FetchSingleUser)
		//v1.PUT("/:id", UpdateUser)
		//v1.DELETE("/:id", DeleteUser)
	}
	router.Run()
}

