package handler

import (
	"log"
	"net/http"
	"telkomGin/config"
	"telkomGin/models"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome",
	})
	return
}

func ShowAllUser(c *gin.Context) {
	con := config.Connect()
	var users []models.User

	err := con.Model(&users).Select()

	if err != nil {
		log.Printf("Gagal Mendapatkan User: %v/n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All User Recorded",
		"data":    users,
	})
	return

}

func SingleUser(c *gin.Context) {
	id := c.Param("id")
	user := &models.User{ID: id}
	err := config.Connect().Select(user)

	if err != nil {
		log.Printf("Gagagl Mendapatkan User")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "found this",
		"data":    user,
	})
	return
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	users := &models.User{ID: id}
	err := config.Connect().Delete(users)
	if err != nil {
		log.Printf("Error, Gagal menghapus user : %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil Menghapus User",
	})
	return
}
