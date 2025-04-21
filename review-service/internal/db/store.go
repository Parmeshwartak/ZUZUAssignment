package db

import (
	"review-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(dsn string) (*Store, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Review{})
	return &Store{DB: db}, nil
}

func (s *Store) InsertReview(r models.Review) error {
	return s.DB.Create(&r).Error
}

