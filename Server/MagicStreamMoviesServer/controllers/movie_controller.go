package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/iaylee/MovieRec/Server/MagicStreamMoviesServer/database"
	"github.com/iaylee/MovieRec/Server/MagicStreamMoviesServer/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

var validate = validator.New()

//query movies collection and return that data to a client via HTTP (w/ gingonic)
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context){
		//c.JSON(200, gin.H{"message":"List of movies"})

		// housekeeping: no matter what happens in this method, our query will timeout and clear up any hanging resources
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies []models.Movie
		cursor, err := movieCollection.Find(ctx, bson.M{})
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies."})
		}
		
		c.JSON(http.StatusOK, movies)
	}
}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		movieID := c.Param("imdb_id")

		if movieID == ""{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Movie ID is required"})
			return
		}
		var movie models.Movie

		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error":"Movie not found"})
		}

		c.JSON(http.StatusOK, movie)

	}
}

func AddMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		//context.Background is not cancelable within itself so we wrap in timeout context
		// cancel is a function that manually cancels the context before the timeout
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		//need to call cancel to prevent context leaks
		defer cancel()

		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
			return
		}
		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Validation failed", "details": err.Error()})
			return
		}

		//movie is the go struct (map) that will be converted into bson document
		// and inserted into the mongoDB movies collection
		result, err := movieCollection.InsertOne(ctx, movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to add movie"})
			return
		}

		//send result back to client
		c.JSON(http.StatusCreated, result)

	}
}