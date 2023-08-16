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

type Pduc_request struct {
	gorm.Model
	Productname string `json:"productname"`
	Detail      string `json:"detail"`
	Price       string `json:"price"`
	Unit        string `json:"unit"`
}

type Pduc_response struct {
	Pduc_request
	ID uint `json:"id"`
}

var produc_sql *gorm.DB = config.ConnectDB()

func ListProductAll(c *gin.Context) {

	var getuser []models.Products
	err := produc_sql.Find(&getuser)

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

func ListProductByid(c *gin.Context) {

}

func CreateProduct(c *gin.Context) {

	var data Pduc_request
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	//Matching todo models struct with todo request struct
	users := models.Products{}
	users.Productname = data.Productname
	users.Detail = data.Detail
	users.Price = data.Price
	users.Unit = data.Unit

	result := produc_sql.Create(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Somethong went wrong"})

		return
	}
	//Matching result to create response
	var response Pduc_response
	response.ID = users.ID
	response.Productname = users.Productname
	response.Detail = users.Detail
	response.Price = users.Price
	response.Unit = users.Unit

	c.JSON(http.StatusCreated, response)
}

func UpdateProduct(c *gin.Context) {
	var data Pduc_request
	reqParamID := c.Param("userid")
	userid := cast.ToUint(reqParamID)
	if err := c.BindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users := models.Products{}
	usersByid := produc_sql.Where("id = ?", userid).First(&users)
	if usersByid.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	users.Productname = data.Productname
	users.Detail = data.Detail
	users.Price = data.Price

	result := produc_sql.Save(&users)
	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response Pduc_response
	response.ID = users.ID
	response.Productname = users.Productname
	response.Detail = users.Detail
	response.Price = users.Price

	c.JSON(http.StatusCreated, response)

}

func DeteleProduct(c *gin.Context) {

	users := models.Products{}

	reqParamId := c.Param("userid")
	userid := cast.ToUint(reqParamId)
	delele := produc_sql.Where("id = ?", userid).Unscoped().Delete(&users)
	fmt.Println(delele)

	if delele.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delete Error"})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    userid,
	})

}
