package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type userRepo struct {
	db *sql.DB
}

// CreateUser implements interfaces.UserRepository
func (c *userRepo) CreateUser(ctx context.Context, user domain.User) (int, error) {
	var id int

	query := `INSERT INTO users (phone,email,password,user_type,verification,status) 
				VALUES($1,$2,$3,$4,$5,$6) RETURNING id_user;`

	err := c.db.QueryRow(query,
		user.Phone,
		user.Email,
		user.Password,
		user.UserType,
		user.Verification,
		user.Status,
	).Scan(
		&id,
	)

	return id, err
}

// FindUserWithEmail implements interfaces.UserRepository
func (c *userRepo) FindUserWithEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	query := `SELECT id_user, phone, email, password, user_type, verification, status from users WHERE email=$1;`

	err := c.db.QueryRow(query,
		email).Scan(
		&user.IdUser,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.UserType,
		&user.Verification,
		&user.Status,
	)
	if err != nil && err == sql.ErrNoRows {
		return user, errors.New("there is no user")
	}

	return user, err
}

// FindUserWithNumber implements interfaces.UserRepository
func (c *userRepo) FindUserWithNumber(ctx context.Context, phoneNumber string) (domain.User, error) {
	var user domain.User
	query := `SELECT id_user, phone, email, password, user_type, verification, status from users WHERE phone=$1;`

	err := c.db.QueryRow(query,
		phoneNumber).Scan(
		&user.IdUser,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.UserType,
		&user.Verification,
		&user.Status,
	)
	if err != nil && err == sql.ErrNoRows {
		return user, errors.New("there is no user")
	}

	return user, err
}

func NewUserRepo(db *sql.DB) interfaces.UserRepository {
	return &userRepo{
		db: db,
	}
}
