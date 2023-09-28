package repository

import (
	"context"
	"ottoDigital/model"
)

func (r *repository) CreateUser(ctx context.Context, user model.User) (uint64, error) {
	query := "INSERT INTO users (name, email, password) values (?,?,?)"

	result, err := r.database.Master.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		return 0, nil
	}

	returnID, _ := result.LastInsertId()

	return uint64(returnID), err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT * FROM users WHERE email = ?"

	result, err := r.database.Master.QueryContext(ctx, query, email)

	if err != nil {
		return nil, err
	}

	var user = &model.User{}

	if result.Next() {
		result.Scan(user.ToPersistent()...)
	}

	return user, nil
}

func (r *repository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	query := "SELECT * FROM users"

	result, err := r.database.Master.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0)

	for result.Next() {
		var user = &model.User{}
		result.Scan(user.ToPersistent()...)
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetUserByID(ctx context.Context, userId int) (*model.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	result, err := r.database.Master.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, err
	}
	var user = &model.User{}

	if result.Next() {
		result.Scan(user.ToPersistent()...)
	}

	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, userId int, user model.User) (bool, error) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"

	result, err := r.database.Master.ExecContext(ctx, query, user.Name, user.Email, userId)

	if err != nil {
		return false, err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *repository) DeleteUser(ctx context.Context, userId int) (bool, error) {
	query := "DELETE FROM users WHERE id = ?"

	result, err := r.database.Master.ExecContext(ctx, query, userId)

	if err != nil {
		return false, err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
