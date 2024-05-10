package entity

import "time"

// Represents user in database
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	Active    bool      `json:"active"`
}
