package controllers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return &UserRepo{Db: db}
}
func (r *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	err = models.CreateUser(r.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, user)

}

func (r *UserRepo) GetUsers(context *gin.Context) {
	var user []models.User
	err := models.GetUsers(r.Db, &user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println(user)
	context.JSON(http.StatusOK, user)

}
func (r *UserRepo) GetUser(c *gin.Context) {
	atoi, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := models.GetUser(r.Db, &user, atoi)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}
	c.JSON(http.StatusOK, user)

}
func (r *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(r.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return

		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	_ = c.Bind(&user)
	err = models.UpdateUser(r.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (r *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(r.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	return

}
