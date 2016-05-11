package main

import "time"

// Event イベント
type Event struct {
	ID          int        `json:"id"`
	EventName   string     `json:"event_name" validate:"required"`
	RoomName    string     `json:"room_name"`
	Description string     `json:"description"`
	Items       string     `json:"items" validate:"required"`
	Major       int        `json:"major"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// User ユーザ
type User struct {
	ID          string     `json:"id"`
	UserName    string     `json:"user_name" validate:"required"`
	Profile     string     `json:"profile"`
	Items       string     `json:"items" validate:"required"`
	Major       int        `json:"major"`
	Image       string     `json:"image"`
	ImageHeader string     `json:"image_header"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
