package storage

import (
	"database/sql"
	"github.com/google/uuid"
	"add/models"
	"log"
)

func InsertMovie(db *sql.DB, movie models.Movie) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO movie (
			id,
			title,
			duration,
			description
		) VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(query,
		id,
		movie.Title,
		movie.Duration,
		movie.Description,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdMovie(db *sql.DB, id string) (models.Movie, error) {

	var (
		movie models.Movie
	)

	query := `
		SELECT
			id,
			title,
			TO_CHAR(duration, 'HH24:MI:SS'),
			description
		FROM movie WHERE id = $1
	`

	err := db.QueryRow(query, id).Scan(
		&movie.Id,
		&movie.Title,
		&movie.Duration,
		&movie.Description,
	)

	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func GetByIdMovie1(db *sql.DB) (models.Movie, error) {

	var (
		movie models.Movie
	)

	query := `
		SELECT
			id,
			title,
			TO_CHAR(duration, 'HH24:MI:SS'),
			description
		FROM movie 
	`

	row,err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	// defer row.Close()
 err = row.Scan(
	&movie.Id,
	&movie.Title,
	&movie.Duration,
	&movie.Description,
)
log.Println(movie)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}
