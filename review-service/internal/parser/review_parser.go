package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"review-service/internal/logger"
	"review-service/models"
	"time"
)

type rawReview struct {
	HotelID   int    `json:"hotelId"`
	Platform  string `json:"platform"`
	HotelName string `json:"hotelName"`
	Comment   struct {
		HotelReviewID  int64   `json:"hotelReviewId"`
		ProviderID     int     `json:"providerId"`
		Rating         float64 `json:"rating"`
		ReviewComments string  `json:"reviewComments"`
		ReviewDate     string  `json:"reviewDate"`
	} `json:"comment"`
}

func ParseJL(data []byte) []models.Review {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var reviews []models.Review

	for scanner.Scan() {
		var raw rawReview
		if err := json.Unmarshal(scanner.Bytes(), &raw); err != nil {
			logger.Warn("Invalid JSON line skipped")
			continue
		}
		// Validate hotelId is present and an integer
		if raw.HotelID == 0 {
			logger.Warn("Invalid review: missing or zero hotelId")
			continue
		}
		if raw.HotelID == 0 || raw.Platform == "" || raw.Comment.HotelReviewID == 0 || raw.Comment.ReviewDate == "" {
			logger.Warn("Missing required fields in review, skipping")
			continue
		}

		parsedTime, err := time.Parse(time.RFC3339, raw.Comment.ReviewDate)
		if err != nil {
			logger.Warn("Invalid date format, skipping")
			continue
		}

		review := models.Review{
			HotelID:        raw.HotelID,
			Platform:       raw.Platform,
			HotelName:      raw.HotelName,
			HotelReviewID:  raw.Comment.HotelReviewID,
			ProviderID:     raw.Comment.ProviderID,
			Rating:         raw.Comment.Rating,
			ReviewComments: raw.Comment.ReviewComments,
			ReviewDate:     parsedTime,
			Language:       "en",
		}
		reviews = append(reviews, review)
	}

	return reviews
}
