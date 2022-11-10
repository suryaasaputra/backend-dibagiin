package models

import "time"

type Donation struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	PhotoUrl     string     `json:"photo_url"`
	Location     string     `json:"location"`
	Status       string     `json:"status"`
	DonationType string     `json:"donation_type"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
