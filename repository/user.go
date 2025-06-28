package repository

import (
	"database/sql"
	"testMod/dto"
)

type UserRepository interface {
	Create(userCreateRepository dto.UserCreateRepository) error
	Update(userUpdateRepository dto.UserUpdateRepository) error
	Delete(id uint) error
	List() ([]dto.UserGetRepository, error)
}

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return userRepository{
		conn: conn,
	}
}
func (userRepo userRepository) Create(userCreateRepository dto.UserCreateRepository) error {
	_, err := userRepo.conn.Exec(`insert into users (username,email,age) values ($1,$2,$3)`,
		userCreateRepository.Name,
		userCreateRepository.Email,
		userCreateRepository.Age)
	return err
}

func (userRepo userRepository) Update(userUpdateRepository dto.UserUpdateRepository) error {
	_, err := userRepo.conn.Exec(`UPDATE users SET username=$1,email=$2,age=$3 WHERE id=$4`,
		userUpdateRepository.Name,
		 userUpdateRepository.Email,
		  userUpdateRepository.Age,
		   userUpdateRepository.Id)
	return err
}
func (userRepo userRepository) Delete(id uint) error {
	_, err := userRepo.conn.Exec(`DELETE FROM USERS WHERE id=$1`, id)
	return err
}

func (userRepo userRepository) List() ([]dto.UserGetRepository, error) {

	var users []dto.UserGetRepository
	// var users []models.User

	rows, err := userRepo.conn.Query(`select id,username,email,age from users`)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user dto.UserGetRepository
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
