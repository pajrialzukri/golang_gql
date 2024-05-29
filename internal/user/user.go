package user

import (
	"backend/graph/model"
	db "backend/pkg/database"
	"fmt"
	"log"
)

type User model.User

func (user *User) Get(id *string) (*model.User, error) {
	var ret model.User
	var err error

	q := `
	SELECT
		id,
		email,
		phone,
		otp,
		source
	FROM users WHERE id = ?
	`
	row := db.Handle.QueryRow(q, id)
	db.LogSQL(q, id)
	err = row.Scan(
		&ret.ID,
		&ret.Email,
		&ret.Phone,
		&ret.Otp,
		&ret.Source,
	)

	if err != nil {
		log.Printf("Error select person: %v", err)
		return nil, err
	}

	return &ret, err
}

func (user *User) Create(data model.UserInput) (*model.User, error) {

	q := `
	INSERT INTO users
(id, email, phone, otp, source)
VALUES(?, ?, ?, ?, ?);
	`
	_, err := db.Handle.Exec(
		q,
		data.ID,
		data.Email,
		data.Phone,
		data.Otp,
		data.Source,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	obj, err := user.Get(&data.ID)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
