package controllers

import (
	"fmt"
	"goapi/config"
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Response struct {
	Request
	ID uint `json:"id"`
}

var user_sql *gorm.DB = config.ConnectDB()

func ListUserAll(c *gin.Context) {

	var getuser []models.Users
	err := user_sql.Find(&getuser)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    getuser,
	})

}

func ListUserByid(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

	var data Request
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	//Matching todo models struct with todo request struct
	users := models.Users{}
	users.Name = data.Name
	users.Lastname = data.Lastname
	users.Address = data.Address
	users.Phone = data.Phone
	users.Username = data.Username
	users.Email = data.Email
	users.Password = data.Password
	users.Role = data.Role

	result := user_sql.Create(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Somethong went wrong"})

		return
	}
	//Matching result to create response
	var response Response
	response.ID = users.ID
	response.Name = users.Name
	response.Lastname = users.Lastname
	response.Address = users.Address
	response.Phone = users.Phone
	response.Username = users.Username
	response.Email = users.Email
	response.Password = users.Password
	response.Role = users.Role
	c.JSON(http.StatusCreated, response)
}

func UpdateUser(c *gin.Context) {
	var data Request
	reqParamID := c.Param("userid")
	userid := cast.ToUint(reqParamID)
	if err := c.BindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users1 := models.Users{}
	usersByid := user_sql.Where("id = ?", userid).First(&users1)
	if usersByid.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	users1.Name = data.Name
	users1.Lastname = data.Lastname
	users1.Address = data.Address
	users1.Phone = data.Phone
	users1.Username = data.Username
	users1.Email = data.Email
	users1.Password = data.Password
	users1.Role = data.Role
	result := user_sql.Save(&users1)
	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response Response
	response.ID = users1.ID
	response.Name = users1.Name
	response.Lastname = users1.Lastname
	response.Address = users1.Address
	response.Phone = users1.Phone
	response.Username = users1.Username
	response.Email = users1.Email
	response.Password = users1.Password
	response.Role = users1.Role

	c.JSON(http.StatusCreated, response)

}

func DeteleUser(c *gin.Context) {

	users := models.Users{}

	reqParamId := c.Param("userid")
	userid := cast.ToUint(reqParamId)
	delele := user_sql.Where("id = ?", userid).Unscoped().Delete(&users)
	fmt.Println(delele)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    userid,
	})

}
