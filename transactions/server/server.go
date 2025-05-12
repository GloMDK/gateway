package server

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func New() (*Server, error) {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{DSN: "host=db user=postgres password=postgres dbname=transactions port=5432"},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	return &Server{db: db}, nil
}
