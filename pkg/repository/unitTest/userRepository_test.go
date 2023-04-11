package unittest

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_FindUserWithNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)

	rows := sqlmock.NewRows([]string{"id_user", "user_name", "phone", "email", "password", "user_type", "verification", "status", "profilephoto"}).
		AddRow(0, "user_name", "1", "email", "password", "user_type", true, "status", "profilephoto")

	tests := []struct {
		name          string
		phoneNumber   string
		mockQueryFunc func()
		expectedUser  domain.User
		expectedErr   error
	}{
		{
			name:        "there is no user in database",
			phoneNumber: "1",
			mockQueryFunc: func() {
				mock.ExpectQuery("SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE phone=\\$1;").
					WithArgs("1").
					WillReturnError(sql.ErrNoRows)
			},
			expectedUser: domain.User{},
			expectedErr:  errors.New("there is no user"),
		},
		{
			name:        "found user from database",
			phoneNumber: "1",
			mockQueryFunc: func() {
				mock.ExpectQuery("SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE phone=\\$1;").
					WithArgs("1").
					WillReturnRows(rows)
			},
			expectedUser: domain.User{Phone: "1"},
			expectedErr:  nil,
		},
		{
			name:        "DB error",
			phoneNumber: "1",
			mockQueryFunc: func() {
				mock.ExpectQuery("SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE phone=\\$1;").
					WithArgs("1").
					WillReturnError(errors.New("DB error"))
			},
			expectedErr: errors.New("DB error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQueryFunc()
			ctx := context.Background()

			actualUser, actualerr := userRepo.FindUserWithNumber(ctx, tt.phoneNumber)

			assert.Equal(t, tt.expectedErr, actualerr)

			assert.Equal(t, tt.expectedUser.Phone, actualUser.Phone)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestUserRepo_FindUserWithemail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)

	rows := sqlmock.NewRows([]string{"id_user", "user_name", "phone", "email", "password", "user_type", "verification", "status", "profilephoto"}).
		AddRow(0, "user_name", "phone", "jon@gmail.com", "password", "user_type", true, "status", "profilephoto")

	tests := []struct {
		name          string
		email         string
		mockQueryFunc func()
		expectedUser  domain.User
		expectedErr   error
	}{
		{
			name:  "there is no user in database",
			email: "jon@gmail.com",
			mockQueryFunc: func() {
				mock.ExpectQuery("SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE email=\\$1;").
					WithArgs("jon@gmail.com").
					WillReturnError(sql.ErrNoRows)
			},
			expectedUser: domain.User{},
			expectedErr:  errors.New("there is no user"),
		},
		{
			name:  "found user from database",
			email: "jon@gmail.com",
			mockQueryFunc: func() {
				mock.ExpectQuery("SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE email=\\$1;").
					WithArgs("jon@gmail.com").
					WillReturnRows(rows)
			},
			expectedUser: domain.User{Email: "jon@gmail.com"},
			expectedErr:  nil,
		},
		{
			name:  "DB error",
			email: "jon@gmail.com",
			mockQueryFunc: func() {
				mock.ExpectQuery("SELECT id_user, user_name, phone, email, password, user_type, verification, status,profilephoto from users WHERE email=\\$1;").
					WithArgs("jon@gmail.com").
					WillReturnError(errors.New("DB error"))
			},
			expectedErr: errors.New("DB error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQueryFunc()
			ctx := context.Background()

			actualUser, actualerr := userRepo.FindUserWithEmail(ctx, tt.email)

			assert.Equal(t, tt.expectedErr, actualerr)

			assert.Equal(t, tt.expectedUser.Email, actualUser.Email)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
