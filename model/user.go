package model

import "ottoDigital/dto"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdateAt  string
}

func (u *User) ToPersistent() []interface{} {
	return []interface{}{
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdateAt,
	}
}

func (u *User) ToDtoResponse() *dto.UserResponse {
	return &dto.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
