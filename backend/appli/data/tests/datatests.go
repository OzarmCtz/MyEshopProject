package datatest

import (
	"database/sql"
	"time"

	"github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var CreateUserParams = mysql.CreateUserParams{
	UEmail:        "johndoe@gmail.com",
	URegisterDate: time.Now(),
	UIsDisabled:   false,
}

var CreateItemParams = mysql.CreateItemParams{
	ITitle:       "Item 1",
	IDescription: mysql.NullString{NullString: sql.NullString{String: "Description for Item 1", Valid: true}},
	IPrice:       "0.0015",
	IQuantity:    mysql.NullInt32{NullInt32: sql.NullInt32{Int32: 1, Valid: true}},
	IPictureUrl:  mysql.NullString{NullString: sql.NullString{String: " https://domain.com/images/items1.png", Valid: true}},
	IFilePath:    mysql.NullString{NullString: sql.NullString{String: "https://domain.com/files/items1.pdf", Valid: true}},
	IIsDisabled:  false,
	IReleaseDate: time.Now(),
}
