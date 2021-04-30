package Middlewares

import "github.com/gin-gonic/gin"

func RespondWithError(code string, c *gin.Context) {
	var message string
	var rescode int
	switch code {

	case "InvalidFormat":
		message = "The card number should only contain digits"
		rescode = 400
	case "InvalidLength":
		message = "The card number must contain minimum of 13 digits"
		rescode = 400
	case "Invalid Card Number":
		message = "Invalid Card Number!"
		rescode = 400
	case "InternalServerError":
		message = "Internal Server Error"
		rescode = 500

	}

	response := map[string]string{"error": message}

	c.JSON(rescode, response)
	c.Abort()
}
