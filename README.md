# API to validate a Debit/Credit Card using Luhn Algorithm

**Requirements:**
Go,
GIN 

**Installation**


1.To install Gin package, you need to install Go and set your Go workspace first.

After installing Go (version 1.13+ is required), then you can use the below Go command to install Gin.

`$ go get  github.com/gin-gonic/gin`

Import it in your code:

`import "github.com/gin-gonic/gin"`

(Optional) Import net/http. This is required for example if using constants such as http.StatusOK.

`import "net/http"`

**Check dbconfig file for database info**



```
$ go run validate.go
```
Do a get call to "_url_"/validate/"_card digits_" (If running on local host use localhost:"_port number_" for URL)

**This is a Basic Auth API call, for API key generation please check the repository : https://github.com/venkatkrishna07/BasicAuth-using-Gin-Gonic for generating API keys.**
