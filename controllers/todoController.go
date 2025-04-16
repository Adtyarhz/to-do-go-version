package controllers

import (
	"go-lang-p1/config"
	"go-lang-p1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    var todos []models.Todo
    config.DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var input models.Todo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&input)
    c.JSON(http.StatusCreated, input)
}

func GetTodoByID(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
        return
    }

    var input models.Todo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Model(&todo).Updates(input)
    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
        return
    }
    config.DB.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}
