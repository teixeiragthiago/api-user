package entity

import "time"

// Represents user in database
type User struct {
	ID        uint      `gorm:"primary_key auto_increment"`
	Name      string    `gorm:"type:varchar(128)"`
	Email     string    `gorm:"type:varchar(128)"`
	Nick      string    `gorm:"type:varchar(128)"`
	Password  []byte    `gorm:"type:varbinary(128)"`
	CreatedAt time.Time `gorm:"not null autoCreateTime:milli"`
	Active    bool      `gorm:"not null default:true"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
}
