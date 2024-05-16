package entity

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey"`
	Content     string    `gorm:"type:varchar(128)"`
	CreatedDate time.Time `gorm:"not null autoCreateTime:milli"`
	IsDeleted   bool      `gorm:"not null default:false"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}
