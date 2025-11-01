package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"` //uniquely identifies each genre within database
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required, min=2, max=100"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"` //each RankingName will have an associated value (quantifiable value)
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

type Movie struct {
	ID          bson.ObjectID `bson:"_id" json:"_id"`                             //used to uniquely identify a movie document in database, bson resides in the MongoDB Go driver
	ImdbID      string        `bson:"imdb_id" json:"imdb_id" validate:"required"` //used to uniquely identify a movie
	Title       string        `bson:"title" json:"title" validate:"required, min=2, max=500"`
	PosterPath  string        `bson:"poster_path" json:"poster_path" validate:"required, url"` //points to URL for poster image
	YoutubeID   string        `bson:"youtube_id" json:"youtube_id" validate:"required"`        //Youtube video ID for the movie's trailer
	Genre       []Genre       `bson:"genre" json:"genre" validate:"required, dive"`            //a movie can have multiple genres
	AdminReview string        `bson:"admin_review" json:"admin_review" validate:"required"`
	Ranking     Ranking       `bson:"ranking" json:"ranking" validate:"required"`
}
