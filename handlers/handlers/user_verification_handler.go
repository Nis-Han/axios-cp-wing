package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/constants"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/models"
)

func (api *Api) generateAndSendOTPViaEmail(c *gin.Context) {
	var user database.User = c.MustGet("userData").(database.User)

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

func (api *Api) verifyUserFromOTP(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)

	verificationParams := models.OTPVerificationParamSchema{}

	err := decoder.Decode(&verificationParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Ill-formatted request body")
		c.Abort()
		return
	}

	var user database.User = c.MustGet("userData").(database.User)

	verificationData, err := api.DB.GetUserVerificationEntryFromUserID(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database Error"})
		c.Abort()
		return
	}

	if verificationData.ValidTill.Before(time.Now()) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Outdated Verification Key. Create a new one!"})
		c.Abort()
		return
	}

	if verificationData.VerificationKey == verificationParams.VerificationKey {
		user, _ = api.DB.SetUserVerificationTrue(c.Request.Context(), verificationData.UserID)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong Verification Key"})
		c.Abort()
		return
	}

	if !user.VerifiedUser {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something Went Wrong, failed to verify user"})
		c.Abort()
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Successfully verified user"})
}
