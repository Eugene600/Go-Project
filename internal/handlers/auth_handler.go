package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/Eugene600/Go-Project/internal/database"
	"github.com/Eugene600/Go-Project/internal/dtos"
	"github.com/Eugene600/Go-Project/internal/models"
	"github.com/Eugene600/Go-Project/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *gin.Context) {
	var reqData dtos.SignUserRequest
	var pgErr *pgconn.PgError

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		log.Printf("Error from bad request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter all required fields",
		})
		return
	}

	// hash password from request
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error while hashing password: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	user := models.User{
		FirstName:    reqData.FirstName,
		LastName:     reqData.LastName,
		DateOfBirth:  reqData.DateOfBirth,
		UserName:     reqData.UserName,
		PasswordHash: string(hashedPassword),
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

	accessToken, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		log.Printf("Error while generating access token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	refreshToken, tokenExpiryDate, err := utils.GenerateRefreshToken(user.Id)
	if err != nil {
		log.Printf("Error while generating refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	//Hash refresh token
	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error while hashing refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	refresh := models.RefreshToken{
		UserId:    user.Id,
		TokenHash: string(hashedRefreshToken),
		ExpiresAt: tokenExpiryDate,
	}

	if err := refresh.CreateRefreshToken(tx, c); err != nil {
		log.Printf("An Error occured while creating refresh token in db, %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
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

	response := dtos.AuthResponse{
		Message:      "User Created Successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, response)
}

func LoginUser(c *gin.Context) {
	var reqData dtos.LoginRequest

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		log.Printf("Error from bad request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter all required fields",
		})
		return
	}

	user := models.User{}

	tx, err := database.DB.BeginTx(c, nil)
	if err != nil {
		log.Printf("Error connecting to DB: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	if err := user.GetUserByUsername(tx, c, reqData.UserName); err != nil {
		log.Printf("Error retrieving user with username of %v: %v", reqData.UserName, err)

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Wrong username or password",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again.",
		})
		return
	}

	// compare hashed password and password
	if err := utils.ComparePassword(user.PasswordHash, reqData.Password); err != nil {
		log.Printf("Error while comparing password: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong username or password",
		})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		log.Printf("Error while generating access token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	refreshToken, tokenExpiryDate, err := utils.GenerateRefreshToken(user.Id)
	if err != nil {
		log.Printf("Error while generating refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error while hashing refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	refresh := models.RefreshToken{
		UserId:    user.Id,
		TokenHash: string(hashedRefreshToken),
		ExpiresAt: tokenExpiryDate,
	}

	if err := refresh.CreateRefreshToken(tx, c); err != nil {
		log.Printf("An Error occured while creating refresh token in db, %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	defer tx.Rollback()

	if err := tx.Commit(); err != nil {
		log.Printf("An Error occured commiting create refresh token transaction, %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong. Please try again",
		})
		return
	}

	response := dtos.AuthResponse{
		Message:      "Login Successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
