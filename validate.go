package main

import (
	"Luhn/Controllers"
	"Luhn/Middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(Middlewares.JSONAppErrorReporter())
	Authorized := r.Group("/", Middlewares.BasicAuth())
	Authorized.GET("/validate/:digits", Controllers.Valid)
	Authorized.GET("/validate/", Controllers.Check)
	r.Run()

}
