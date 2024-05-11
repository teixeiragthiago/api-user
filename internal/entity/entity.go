package entity

import "time"

// Represents user in database
type User struct {
	ID        uint      `gorm:"primary_key auto_increment"`
	Name      string    `gorm:"not null size:50"`
	Nick      string    `gorm:"not null size:15"`
	Password  string    `gorm:"not null size:12"`
	CreatedAt time.Time `gorm:"not null autoCreateTime:milli"`
	Active    bool      `gorm:"not null default:true"`
}
