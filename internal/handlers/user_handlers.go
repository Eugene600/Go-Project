package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/Eugene600/Go-Project/internal/database"
	"github.com/Eugene600/Go-Project/internal/dtos"
	"github.com/Eugene600/Go-Project/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateUser(c *gin.Context) {
	var reqData dtos.CreateUserRequest
	var pgErr *pgconn.PgError

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		log.Printf("Error from bad request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter all required fields",
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

	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error connecting to DB: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	if err := user.CreateUser(tx, c); err != nil {
		
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				log.Printf("An Error occured creating user, %s", err.Error())
				c.JSON(http.StatusConflict, gin.H{
					"error": "Username already exists.",
				})
			default:
				log.Printf("An Error occured creating user, %s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Something went wrong. Please try again",
				})
			}
		}
		return
	}

	defer tx.Rollback()

	if err := tx.Commit(); err != nil {
		log.Printf("An Error occured commiting create user transaction, %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created Successfully",
	})

}
