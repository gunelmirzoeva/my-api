package main

import (
	"net/http"
	//"sort"
	"strconv"
	"github.com/gin-gonic/gin"
)
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Deleted bool   `json:"deleted"`
}

var users []User
var nextID = 1


func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users", createUser)
		api.POST("/users/:id/login", loginUser)
		/*api.PUT("/users/:id/disable", disableUser)
		api.DELETE("/users/:id", deleteUser)
		api.GET("/users", listUsers)*/
	}

	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	newUser.ID = nextID
	nextID ++
	newUser.Status = "Pending"
	newUser.Deleted = false
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func loginUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid ID"})
		return
	}

	for i, u := range users {
		if u.ID == id && !u.Deleted {
			if u.Status == "Pending" {
				users[i].Status = "Active"
			}
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}