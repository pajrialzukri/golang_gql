package user

import (
	"backend/graph/model"
	db "backend/pkg/database"
	"fmt"
	"log"
	"strconv"
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
(email, phone, otp, source)
VALUES( ?, ?, ?, ?);
	`
	create, err := db.Handle.Exec(
		q,
		data.Email,
		data.Phone,
		data.Otp,
		data.Source,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(create)
	lastInsertID, err := create.LastInsertId()
	str := strconv.FormatInt(lastInsertID, 10)
	fmt.Println(str)

	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	obj, err := user.Get(&str)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (user *User) Update(id *string, data model.UserInput) (*model.User, error) {
	q := `
	UPDATE users
	SET email = ?, phone = ?, otp = ?, source = ?
	WHERE id = ?;
	`
	_, err := db.Handle.Exec(
		q,
		data.Email,
		data.Phone,
		data.Otp,
		data.Source,
		id,
	)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	// Retrieve the updated user data
	updatedUser, err := user.Get(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (User *User) Delete(id *string) (bool, error) {
	q := `
	DELETE FROM users
	WHERE id = ?;
	`
	_, err := db.Handle.Exec(q, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return false, err

	}

	return true, nil
}
