package models

import (
	"fmt"

	"github.com/hassanjawwad12/event-management-system/db"
	"github.com/hassanjawwad12/event-management-system/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	securePwd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	fmt.Println("securePwd", securePwd)
	result, err := stmt.Exec(u.Email, securePwd)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}
