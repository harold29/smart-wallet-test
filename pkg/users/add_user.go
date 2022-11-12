package users

import (
	"harold29/yourkeyswallet/pkg/common/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	FirstName    string    `json:"first_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	PhoneNumber1 string    `json:"phone_number_1" binding:"required"`
	PhoneNumber2 string    `json:"phone_number_2"`
	Gender       string    `json:"gender"`
	Birthday     time.Time `json:"birthday"`
}

func (h handler) AddUser(c *gin.Context) {

	body := AddUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Email = body.Email
	user.PhoneNumber1 = body.PhoneNumber1
	user.PhoneNumber2 = body.PhoneNumber2
	user.Birthday = body.Birthday
	user.Gender = body.Gender

	if result := h.DB.Omit("UserType", "UserRole").Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}
