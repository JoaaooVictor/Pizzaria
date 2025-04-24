package main

import (
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{Id: 1, Nome: "Toscana", Preco: 11.20},
	{Id: 2, Nome: "Marguerita", Preco: 13.20},
	{Id: 3, Nome: "Atum com Queijo", Preco: 15.20},
}

func main() {
	var router = gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzaById)
	router.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza = models.Pizza{}
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func getPizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	for _, p := range pizzas {
		if p.Id == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found"})
}
