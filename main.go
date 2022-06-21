package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type db struct {
	List []Post `json:"list"`
}

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Deleted     bool   `json:"-"`
}

type newPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	var database db
	database.List = append(database.List, Post{
		ID:          0,
		Title:       "Test1",
		Description: "Lorem Ipsum",
		Price:       10,
	})

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		var output db

		for _, e := range database.List {
			log.Println(e)
			if e.Deleted == false {
				output.List = append(output.List, e)
			}
		}

		c.JSON(http.StatusOK, output.List)
	})

	router.POST("/addPost", func(c *gin.Context) {
		var postInput newPost
		var postOutput Post

		if err := c.BindJSON(&postInput); err != nil {
			c.Status(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}

		postOutput.ID = len(database.List)
		postOutput.Title = postInput.Title
		postOutput.Description = postInput.Description
		postOutput.Price = postInput.Price

		database.List = append(database.List, postOutput)

		c.JSON(http.StatusOK, postOutput)
	})

	router.DELETE("/:product_id", func(c *gin.Context) {
		idString := c.Param("product_id")

		id, err := strconv.Atoi(idString)
		if err != nil || id < 0 || id > len(database.List)-1 {
			fmt.Println(err.Error())
			return
		}

		database.List[id].Deleted = true

		c.Status(http.StatusOK)
	})

	router.PUT("/put_test", func(c *gin.Context) {

	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
