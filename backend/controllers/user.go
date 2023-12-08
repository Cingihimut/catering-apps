package controllers

import (
	"net/http"

	"github.com/Cingihimut/catering-apps/config"
	"github.com/gin-gonic/gin"
)

type Users struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct{}

func (controller *User) Get(c *gin.Context) {
	var users []Users
	config.DB.Raw("SELECT * FROM users").Scan(&users)
	c.JSON(http.StatusOK, users)
}

func (controller *User) GetOne(c *gin.Context) {
	var user Users
	id := c.Param("id")
	config.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	c.JSON(http.StatusOK, user)
}

func (controller *User) Create(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	c.JSON(http.StatusCreated, user)
}

func (controller *User) Update(c *gin.Context) {
	var user Users
	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, id)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (controller *User) Delete(c *gin.Context) {
	id := c.Param("id")
	config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
