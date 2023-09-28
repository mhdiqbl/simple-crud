package repository

import (
	"context"
	"ottoDigital/model"
)

type IRepository interface {
	//Login(ctx context.Context, user *model.User) (bool, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*model.User, error)
	CreateUser(ctx context.Context, user model.User) (uint64, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, userId int) (*model.User, error)
	UpdateUser(ctx context.Context, userId int, user model.User) (bool, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)

	CreateTask(ctx context.Context, task model.Task) (uint64, error)
	GetAllTasks(ctx context.Context) ([]*model.Task, error)
	GetTaskByID(ctx context.Context, taskId int) (*model.Task, error)
	UpdateTask(ctx context.Context, taskId int, task model.Task) (bool, error)
	DeleteTask(ctx context.Context, taskId int) (bool, error)
}
