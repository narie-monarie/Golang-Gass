package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/narie-monarie/Models/User"
)

func GetPersons(c *gin.Context) {
	persons, err := models.GetPeople(10)
	checkError(err)

	if persons == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
}

func GetPersonById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get Person" + id + " By ID"})
}
func AddPerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Added a person"})
}
func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Updated a Person" + id})
}
func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Person" + id})
}

func Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "called options"})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
