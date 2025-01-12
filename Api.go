package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	Name        string `json : "name"`
	Rate        int    `json : "rate"`
	Color       string `json : "color"`
	Manufacture int    `json : "manufacture"`
}

// inbuild get request of api
var cars = []car{

	{Name: "lamborgini", Rate: 2500000, Color: "black", Manufacture: 2015},
	{Name: "farari", Rate: 2545600, Color: "white", Manufacture: 2014},
	{Name: "scorpio", Rate: 2500000, Color: "green", Manufacture: 2005},
	{Name: "eco-sport", Rate: 255600, Color: "grey", Manufacture: 2014},
	{Name: "tesla", Rate: 7989000, Color: "yellow", Manufacture: 2013},
	{Name: "tesla", Rate: 7989000, Color: "yellow", Manufacture: 2013},
}

// GET request
func getcars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

// Post request to API
func cretecar(c *gin.Context) {
	var newcar car
	if err := c.BindJSON(&newcar); err != nil {
		return
	}
	cars = append(cars, newcar)
	c.IndentedJSON(http.StatusCreated, newcar)
}
func main() {
	router := gin.Default()
	router.GET("/cars", getcars)
	router.Run("localhost:8080")
}
