package entity

import "time"

// Represents user in database
type User struct {
	ID        uint      `gorm:"primary_key auto_increment"`
	Name      string    `gorm:"not null"`
	Nick      string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null autoCreateTime:milli"`
	Active    bool      `gorm:"not null default:true"`
}
