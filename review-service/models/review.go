package models

import "time"

type Review struct {
	ID             uint      `gorm:"primaryKey"`
	HotelID        int
	Platform       string
	HotelName      string
	HotelReviewID  int64     `gorm:"uniqueIndex"`
	ProviderID     int
	Rating         float64
	ReviewComments string
	ReviewDate     time.Time
	Language       string
}

