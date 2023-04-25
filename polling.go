package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type WordCount struct {
	Word  string `json:"user-input"`
	Count int    `json:"total"`
}

var words = make(map[string]int)

func main() {
	r := gin.Default()
	r.POST("/input", func(c *gin.Context) {
		input := c.PostForm("input")
		addWords(input)
		c.JSON(200, gin.H{
			"message": "Input berhasil ditambahkan",
		})
	})
	r.GET("/result", func(c *gin.Context) {
		var results []WordCount
		for word, count := range words {
			results = append(results, WordCount{
				Word:  word,
				Count: count,
			})
		}
		c.JSON(200, results)
	})
	r.Run(":8080")
}

func addWords(input string) {
	wordsInInput := strings.Fields(strings.ToLower(input))	
	for _, word := range wordsInInput {
		words[word]++
	}
}
