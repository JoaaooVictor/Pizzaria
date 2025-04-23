package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	Id    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	var router = gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []Pizza{
		{Id: 1, Nome: "Toscana", Preco: 11.20},
		{Id: 2, Nome: "Marguerita", Preco: 13.20},
		{Id: 3, Nome: "Atum com Queijo", Preco: 15.20},
	}

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
