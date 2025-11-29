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
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

	//router.Run starts a HTTP server using go standard net HTTP package
	//net HTTP package spawns a new go routine for each incoming HTTP request
	//so everything a client makes a request, it is handled concurrently in its own go routine
	if err := router.Run(":8080"); err != nil{
		fmt.Println("Failed to start server", err)
	}
}
