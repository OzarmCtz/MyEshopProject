package validate

import "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"

func CheckCreateUserParams(params mysql.CreateUserParams) ([]string, error) {
	// TODO CHECK INPUT USER CREATED
	fieldsError := []string{}
	return fieldsError, nil
}

func CheckUpdateUserParams(params mysql.UpdateUserParams) ([]string, error) {
	// TODO CHECK INPUT USER UPDATED
	fieldsError := []string{}
	return fieldsError, nil
}
