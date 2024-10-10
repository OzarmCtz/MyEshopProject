package users

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adt "github.com/OzarmCtz/e_shop_backend_v1/appli/data/tests"

	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListUsers() ([]adm.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	users, err := adm.QueriesDb.ListUsers(ctx)
	return users, err
}

func GetUser(userID int32) (adm.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	user, err := adm.QueriesDb.GetUser(ctx, userID)
	return user, err
}

func InsertUser(userParams adm.CreateUserParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateUser(ctx, userParams)
	return res, err
}

func UpdateUser(userParams adm.UpdateUserParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateUser(ctx, userParams)
	return rows, err
}

func DeleteUser(userId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteUser(ctx, userId)
	return rows, err
}

func GetUserByEmail(email string) (adm.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	user, err := adm.QueriesDb.GetUserByEmail(ctx, email)
	return user, err
}

// Utility function
func CreateUser() (adm.User, error) {
	var user adm.User

	resInsertUser, err := InsertUser(adt.CreateUserParams)
	if err != nil {
		return user, nil
	}

	userId, err := resInsertUser.LastInsertId()
	if err != nil {
		return user, err
	}

	user, err = GetUser(int32(userId))
	if err != nil {
		return user, err
	}

	return user, nil
}
