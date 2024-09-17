package main
import (
	"fmt"
	"strconv"
)
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	
	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id");
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error")
		}
		sum := i+3
		c.JSON(200, gin.H{
			"id":sum,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}