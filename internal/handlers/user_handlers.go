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
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateUser(c *gin.Context) {
	var reqData dtos.UserRequest
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

func GetUserByUsername(c *gin.Context) {
	username := c.Query("username")

	if username == "" {
		log.Printf("Error while getting user by username: Username query parameter missing")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username query parameter is required",
		})
		return
	}

	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	defer tx.Rollback()

	var user models.User

	if err := user.GetUserByUsername(tx, c, username); err != nil {
		log.Printf("Error retrieving user: %v", err)

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	response := dtos.MappedUserResponse(user)

	c.JSON(http.StatusOK, response)
}

func GetAllUsers(c *gin.Context) {
	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	defer tx.Rollback()

	users, err := models.GetAllUsers(tx, c)
	if err != nil {
		log.Printf("Error retrieving users: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	response := dtos.MappedUserResponseList(users)

	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	var (
		reqData dtos.UserRequest
		pgErr   *pgconn.PgError
	)

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		log.Printf("Error while updating user: Invalid user id")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id.",
		})
		return
	}

	if err := c.ShouldBindJSON(&reqData); err != nil {
		log.Printf("Error binding request: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter all required fields.",
		})
		return
	}

	user := models.User{
		Id:          id,
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
		log.Printf("Error beginning transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	defer tx.Rollback()

	if err := user.UpdateUser(tx, c); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found.",
			})
			return
		}

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				log.Println("Username already exists.")
				c.JSON(http.StatusConflict, gin.H{
					"error": "Username already exists.",
				})
			default:
				log.Printf("Database error: %v", err)

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Something went wrong. Please try again.",
				})
			}
			return
		}

		log.Printf("Error updating user: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	response := dtos.MappedUserResponse(user)

	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id.",
		})
		return
	}

	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	defer tx.Rollback()

	if err := models.DeleteUser(tx, c, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("User not found when deleting user.")
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found.",
			})
			return
		}

		log.Printf("Error deleting user: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully.",
	})
}

func RecoverDeletedUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id.",
		})
		return
	}

	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	defer tx.Rollback()

	if err := models.RecoverDeletedUser(tx, c, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Deleted user not found.",
			})
			return
		}

		log.Printf("Error recovering user: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User recovered successfully.",
	})
}

func GetDeletedUsers(c *gin.Context) {
	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	defer tx.Rollback()

	users, err := models.GetDeletedUsers(tx, c)
	if err != nil {
		log.Printf("Error retrieving deleted users: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	response := dtos.MappedUserResponseList(users)

	c.JSON(http.StatusOK, response)
}
