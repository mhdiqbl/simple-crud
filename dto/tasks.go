package dto

type TaskRequest struct {
	UserId      int    `json:"user_id" validate:"required,exists:users,id"`
	Tittle      string `json:"tittle" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required,max=50"`
}

type TaskResponse struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdateAt    string `json:"update_at"`
}
