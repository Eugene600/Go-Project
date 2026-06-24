package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Eugene600/Go-Project/internal/database"
	"github.com/Eugene600/Go-Project/internal/dtos"
	"github.com/Eugene600/Go-Project/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var reqData dtos.CreateUserRequest

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		FirstName:   reqData.FirstName,
		LastName:    reqData.LastName,
		DateOfBirth: reqData.DateOfBirth,
		UserName:    reqData.UserName,
	}

	if reqData.MiddleName != nil {
		user.MiddleName = sql.NullString{
			String: *reqData.MiddleName,
			Valid:  true,
		}
	}

	db, err := database.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something happened, please try again",
		})
	}

	tx, err := db.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error connecting to DB: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	if err := user.CreateUser(tx, c); err != nil {
		tx.Rollback()
		log.Printf("An Error occured creating user, %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("An Error occured creating user, %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	c.Status(http.StatusCreated)

}
