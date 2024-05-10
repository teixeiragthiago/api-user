// internal/dto/user_dto.go
package dto

import "time"

type UserDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Active    bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}
