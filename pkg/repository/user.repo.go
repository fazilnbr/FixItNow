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

	query := `INSERT INTO users (user_name,phone,email,password,user_type,verification,status,profilephoto) 
				VALUES($1,$2,$3,$4,$5,$6,$7,$8);`

	err := c.db.QueryRow(query,
		user.UserName,
		user.Phone,
		user.Email,
		user.Password,
		user.UserType,
		user.Verification,
		user.Status,
		user.Profilephoto,
	).Scan(
		&id,
	)

	return id, err
}

// FindUserWithId implements interfaces.UserRepository
func (*userRepo) FindUserWithId(ctx context.Context, id int) (domain.User, error) {
	panic("unimplemented")
}

// FindUserWithNumber implements interfaces.UserRepository
func (c *userRepo) FindUserWithNumber(ctx context.Context, phoneNumber string) (domain.User, error) {
	var user domain.User
	query := `SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE phone=$1;`

	err := c.db.QueryRow(query,
		phoneNumber).Scan(
		&user.IdUser,
		&user.UserName,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.UserType,
		&user.Verification,
		&user.Status,
		&user.Profilephoto,
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
