package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const divider = "-&|@-"

var db = make(map[string]map[string]string)

func register(c *gin.Context) {
	firstname := c.PostForm("firstname")
	if firstname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "firtname is required",
		})
		return
	}

	lastname := c.PostForm("lastname")
	if lastname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "lastname is required",
		})
		return
	}

	dbKey := firstname + divider + lastname
	if _, found := db[dbKey]; found {
		c.JSON(http.StatusConflict, gin.H{
			"message": fmt.Sprintf("%s %s registered before", firstname, lastname),
		})
		return
	}

	job := c.DefaultPostForm("job", "Unknown")
	age := c.DefaultPostForm("age", "18")
	if _, err := strconv.Atoi(age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "age should be integer",
		})
		return
	}

	db[dbKey] = map[string]string{
		"job": job,
		"age": age,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s %s registered successfully", firstname, lastname),
	})
}

func hello(c *gin.Context) {
	firstname := c.Param("firstname")
	lastname := c.Param("lastname")
	dbKey := firstname + divider + lastname
	data, found := db[dbKey]
	if !found {
		c.String(http.StatusNotFound, fmt.Sprintf("%s %s is not registered", firstname, lastname))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("Hello %s %s; Job: %s; Age: %s", firstname, lastname, data["job"], data["age"]))
}
