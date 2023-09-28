package repository

import (
	"context"
	"ottoDigital/model"
)

func (r *repository) CreateTask(ctx context.Context, task model.Task) (uint64, error) {
	query := "INSERT INTO tasks (user_id, tittle, description, status) values (?,?,?,?)"

	result, err := r.database.Master.ExecContext(ctx, query,
		task.UserId,
		task.Tittle,
		task.Description,
		task.Status,
	)

	if err != nil {
		return 0, err
	}

	returnID, _ := result.LastInsertId()

	return uint64(returnID), err
}

func (r *repository) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	query := "SELECT * FROM tasks"

	result, err := r.database.Master.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	tasks := make([]*model.Task, 0)

	for result.Next() {
		var task = &model.Task{}
		result.Scan(task.ToPersistent()...)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *repository) GetTaskByID(ctx context.Context, taskId int) (*model.Task, error) {
	query := "SELECT * FROM tasks WHERE id = ?"

	result, err := r.database.Master.QueryContext(ctx, query, taskId)

	//fmt.Println(err)

	if err != nil {
		return nil, err
	}

	var task = &model.Task{}

	if result.Next() {
		result.Scan(task.ToPersistent()...)
	}

	return task, nil
}

func (r *repository) UpdateTask(ctx context.Context, taskId int, task model.Task) (bool, error) {
	query := "UPDATE tasks SET user_id = ?, tittle = ?, description = ?, status = ? WHERE id = ?"

	result, err := r.database.Master.ExecContext(ctx, query, task.UserId, task.Tittle, task.Description, task.Status, taskId)

	if err != nil {
		return false, err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *repository) DeleteTask(ctx context.Context, taskId int) (bool, error) {
	query := "DELETE FROM tasks WHERE id = ?"

	result, err := r.database.Master.ExecContext(ctx, query, taskId)

	if err != nil {
		return false, err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
