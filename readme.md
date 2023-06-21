# Setup golang by first installing IDE

# Steps
1. https://go.dev/doc/install
2. Install two packages
- Go extention
- Code Runner
3. Use cntrl + shift + P then - Go: Install/Update Tools
4. Go to settings and edit Code runner to run in terminal otherwise run in terminal in file path e.g. go run %path%/helloworld.go


# To setup API you need to get gin to host
## Steps
1. First go to path of package/folder in termninal then run:
- go mod init myproject
- go get -u github.com/gin-gonic/gin
2. Now make sure your .go file is in the same directory then run API.
3. Filter in localhost:8000 to the run and bugfix then via browser/curl. e.g. router.Run("127.0.0.1:8000")

# Simple go API hosting script
package main

// get dependencies
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Start the server
	router.Run("127.0.0.1:8000")
}