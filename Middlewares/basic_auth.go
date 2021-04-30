package Middlewares

import (
	"Luhn/Models"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized, please use Basic Auth as authentication method", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized, please provide a valid API key/secret", c)
			return
		}

		c.Next()
	}

}

func authenticateUser(username, password string) bool {
	var authenticated bool
	db1, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/TestAuth")
	if err != nil {
		panic(err)
	}

	result, err := db1.Query("Select username , pass from apikeys where username = ?  ", username)
	if err != nil {
		return false
	}

	for result.Next() {
		var UserCred Models.User
		err := result.Scan(&UserCred.User, &UserCred.Pass)
		if err != nil {
			panic(err)
		}
		if password == UserCred.Pass {
			authenticated = true
			fmt.Println("Authenticated user : ", username)

		} else {
			authenticated = false

		}
	}
	db1.Close()
	if authenticated {

		return true

	} else {

		return false
	}

}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}
