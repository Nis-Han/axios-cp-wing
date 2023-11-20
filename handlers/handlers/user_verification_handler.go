package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/constants"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

func (api *Api) generateAndSendOTPViaEmail(c *gin.Context) {
	userData := c.MustGet("userData")
	user, ok := userData.(database.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		c.Abort()
		return
	}

	verificationData, err := api.DB.CreateUserVerification(c.Request.Context(), user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database Error"})
		c.Abort()
		return
	}

	err = api.EmailClient.SendEmail(user.Email, constants.OTPVerificationUsingMail(user.FirstName, verificationData.VerificationKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Email Client Error"})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, "")

}
