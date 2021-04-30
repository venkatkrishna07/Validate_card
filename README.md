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

**Check dbconfig file for database info**

`import "net/http"`

```
$ go run validate.go
```
Do a get call to "_url_"/validate/"_card digits_" (If running on local host use localhost:"_port number_" for URL)
