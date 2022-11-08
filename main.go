package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
}

var data = []User{
	{"1", "Gor", "Manucharyang"},
	{"2", "Lorenzo", "Manucci"},
}

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})
	r.POST("/adduser", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		data = append(data, user)
		c.JSON(http.StatusCreated, user)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, a := range data {
			if a.ID == id {
				data = append(data[:i], data[i+1:]...)
				break
			}
		}
		c.Status(http.StatusNoContent)

	})
	r.Run()
}
