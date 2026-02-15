package main

import "github.com/gin-gonic/gin"

type Car struct {
	ID     int
	Amount float64
	UserID int
}

var Orders []Car
var UserBalance map[int]int

func init() {
	UserBalance = map[int]int{
		1: 100,
		2: 200,
		3: 300,
	}
	Orders = []Car{}
}
func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
