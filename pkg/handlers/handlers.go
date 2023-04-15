package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func CalculateNumberOfSymbols(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	text := queryParams.Get("text")
	if text == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Provide the 'text' param to perform calculations",
		})
		return
	}

	numberOfSymbols := utf8.RuneCountInString(text)
	c.String(http.StatusOK, strconv.Itoa(numberOfSymbols))
}

func CalculateNumberOfWords(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	text := queryParams.Get("text")
	if text == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Provide the 'text' param to perform calculations",
		})
		return
	}

	numberOfWords := len(strings.Split(text, " "))
	c.String(http.StatusOK, strconv.Itoa(numberOfWords))
}

func CalculateReadingTime(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	text := queryParams.Get("text")
	if text == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Provide the 'text' param to perform calculations",
		})
		return
	}

	numberOfWords := len(strings.Split(text, " "))
	readingTime := numberOfWords / 180
	if readingTime == 0 {
		readingTime++
	}
	c.String(http.StatusOK, strconv.Itoa(readingTime))
}
