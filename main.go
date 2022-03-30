package main

import (
	"insider_interview/internal/database"
	"insider_interview/internal/job"
	"insider_interview/internal/repository/concrete"
)

func main() {
	db := concrete.NewMemDB()

	api := database.API{
		Db: db,
	}

	job.BeginGame(api)
}
