package Models

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//This DB connection for other database func , not for authentication
func ConnectDatabase() {
	var c *gin.Context
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/TestAuth")
	if err != nil {
		respondWithError(500, "Server Error!", c)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		respondWithError(500, "Server Error!", c)
		panic(err)
	}

	DB = db

}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
