package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "github.com/iaylee/MovieRec/Server/MagicStreamMoviesServer/controllers"
)

func main() {
	// creates gin router with default middleware
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context){
		c.String(200, "MoviesRec")
	})

	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil{
		fmt.Println("Failed to start servier", err)
	}
}
