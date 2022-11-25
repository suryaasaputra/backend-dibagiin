package models

import (
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type DonationHistory struct {
	ID         string     `json:"id" gorm:"primaryKey;type:varchar"`
	UserID     string     `json:"user_id" gorm:"not null;"`
	DonationID string     `json:"donation_id" gorm:"not null;"`
	User       *User      `json:"-"`
	Donation   *Donation  `json:"-"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (d *DonationHistory) BeforeCreate(tx *gorm.DB) error {
	newId := xid.New().String()
	d.ID = newId
	return nil
}
