package model

import "ottoDigital/dto"

type Task struct {
	ID          int
	UserId      int
	Tittle      string
	Description string
	Status      string
	CreatedAt   string
	UpdateAt    string
}

func (t *Task) ToPersistent() []interface{} {
	return []interface{}{
		&t.ID,
		&t.UserId,
		&t.Tittle,
		&t.Description,
		&t.Status,
		&t.CreatedAt,
		&t.UpdateAt,
	}
}

func (t *Task) ToDtoResponse() *dto.TaskResponse {
	return &dto.TaskResponse{
		ID:          t.ID,
		UserId:      t.UserId,
		Tittle:      t.Tittle,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		UpdateAt:    t.UpdateAt,
	}
}
