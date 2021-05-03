package Controllers

import (
	"Luhn/Middlewares"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ")
	c.JSON(http.StatusBadRequest, "Please provide a card number! ")
}

func Valid(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length,Authorization")
	param := c.Param("digits")

	res := Validate(param)
	if res == "Success" {
		c.JSON(http.StatusOK, "Valid Card Number")
	} else {
		Middlewares.RespondWithError(res, c)
		c.Error(errors.New(res))
	}

}

func Validate(cardno string) string {
	sum := 0
	cardno1 := strings.ReplaceAll(cardno, " ", "")
	r := []rune(cardno1)
	revr := Reverse(r)
	if len(revr) > 13 {
		if checkdigit([]rune(cardno1)) {
			number := strings.Split(revr, "")
			for i := 1; i < len(number); i += 2 {

				num, err := strconv.Atoi(number[i])
				if err != nil {
					//respondWithError(500, "Internal Server Error")
					return "InternalServerError"
				}
				num = num * 2
				if num > 9 {
					num = num - 9
				}
				fnum := strconv.Itoa(num)
				number[i] = fnum

			}
			for i := 0; i < len(number); i++ {

				n, err := strconv.Atoi(number[i])
				if err != nil {
					//respondWithError(500, "Internal Server Error")
					return "InternalServerError"
				}
				sum = sum + n
			}
			if sum%10 == 0 {

				return "Success"
			} else {
				//respondWithError(400, "Invalid Card Number")
				return "Invalid Card Number"
			}

		} else {
			//respondWithError(400, "The card number can only have digits")
			return "InvalidFormat"
		}

	} else {
		//respondWithError(400, "The length of the card number must be of 16 digits!")
		return "InvalidLength"
	}

}

func checkdigit(r []rune) bool {
	counter := 0
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)
	for i := 0; i < len(r); i++ {
		//fmt.Println(r[i])
		if !digitCheck.MatchString(string(r[i])) {
			counter++
		}

	}
	if counter > 0 {

		return false
	}

	return true
}

func Reverse(s []rune) string {
	runes := s
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
