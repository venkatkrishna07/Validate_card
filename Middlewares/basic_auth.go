package Middlewares

import (
	"Luhn/Models"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		//Check if the autorization is Basic Auth and if both key and secret are present
		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized, please use Basic Auth as authentication method", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		//Authenticate the user
		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized, please provide a valid API key/secret", c)
			return
		}

		c.Next()
	}

}

func authenticateUser(key, password string) bool {
	var authenticated bool
	server_secret := "<server_secret>" // dummy value 
	db1, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/Auth") //DB connection
	if err != nil {
		panic(err)
	}
	defer db1.Close()
	result, err := db1.Query("Select hash_value  from api_keys where api_key = ?  ", key)
	if err != nil {
		return false
	}
	new_hash_value := ValidateHmac256(password, server_secret) //Get the HMAC value of the secret
	for result.Next() {
		var UserCred Models.User
		err := result.Scan(&UserCred.Hash_value)
		if err != nil {
			panic(err)
		}

		//Compare the generated HMAC value and the HMAC value stored in db
		if new_hash_value == UserCred.Hash_value {
			authenticated = true
			fmt.Println("Authenticated user : ", key)

		} else {
			authenticated = false

		}
	}
	db1.Close()
	return authenticated

}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

//Function to compute HMAC value
func ValidateHmac256(key_secret string, server_secret string) string {
	key_sec := []byte(server_secret)
	h := hmac.New(sha256.New, key_sec)
	h.Write([]byte(key_secret))
	hash_value := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return hash_value

}
