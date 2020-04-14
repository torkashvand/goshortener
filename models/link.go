package models

import (
	"time"
)

// Model is base model for application models
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Link modle
type Link struct {
	Model
	Address  string
	Shortcut string
}
