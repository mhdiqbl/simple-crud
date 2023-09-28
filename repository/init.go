package repository

import "ottoDigital/database"

type repository struct {
	database *database.DatabaseComponent
}

func NewRepository() IRepository {
	db := database.InitDatabase()

	return &repository{
		database: db,
	}
}
