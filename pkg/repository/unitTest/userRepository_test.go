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

func TestUserRepo_AddProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)

	mockQuery := "INSERT INTO profiles \\(user_id, first_name, last_name, gender, dob, profile_photo\\) VALUES \\(\\$1,\\$2,\\$3,\\$4,\\$5,\\$6\\) RETURNING id_profie;"
	mockUserData := domain.UserData{
		UserId:       0,
		Email:        "",
		FirstName:    "",
		LastName:     "",
		Gender:       "",
		Dob:          "",
		ProfilePhoto: "",
	}

	tests := []struct {
		name          string
		UserData      domain.UserData
		mockQueryFunc func()
		expectedErr   error
	}{
		{
			name:     "test success adding profile",
			UserData: mockUserData,
			mockQueryFunc: func() {
				mock.ExpectQuery(mockQuery).
					WithArgs(mockUserData.UserId, mockUserData.FirstName, mockUserData.LastName, mockUserData.Gender, mockUserData.Dob, mockUserData.ProfilePhoto).
					WillReturnRows(sqlmock.NewRows([]string{"id_profie"}).AddRow(1))
			},
			expectedErr: nil,
		},
		{
			name:     "test there is no user with id",
			UserData: mockUserData,
			mockQueryFunc: func() {
				mock.ExpectQuery(mockQuery).
					WithArgs(mockUserData.UserId, mockUserData.FirstName, mockUserData.LastName, mockUserData.Gender, mockUserData.Dob, mockUserData.ProfilePhoto).
					WillReturnRows(sqlmock.NewRows([]string{"id_profie"}).AddRow(0))
			},
			expectedErr: errors.New("Invalid User"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQueryFunc()
			ctx := context.Background()

			actualerr := userRepo.AddProfile(ctx, tt.UserData)

			assert.Equal(t, tt.expectedErr, actualerr)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestUserRepo_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)

	mockQuery := "INSERT INTO users \\(phone,email,password,user_type,verification,status\\) VALUES\\(\\$1,\\$2,\\$3,\\$4,\\$5,\\$6\\) RETURNING id_user;"
	mockUser := domain.User{
		IdUser:       1,
		Phone:        "",
		Email:        "",
		Password:     "",
		UserType:     "",
		Verification: false,
		Status:       "",
	}

	tests := []struct {
		name          string
		user          domain.User
		mockQueryFunc func()
		expectedId    int
		expectedErr   error
	}{
		{
			name: "test there is any db error ",
			user: mockUser,
			mockQueryFunc: func() {
				mock.ExpectQuery(mockQuery).
					WithArgs(mockUser.Phone, mockUser.Email, mockUser.Password, mockUser.UserType, mockUser.Verification, mockUser.Status).
					WillReturnError(errors.New("db error"))
			},
			expectedId:  0,
			expectedErr: errors.New("db error"),
		},
		{
			name: "test success creating user",
			user: mockUser,
			mockQueryFunc: func() {
				mock.ExpectQuery(mockQuery).
					WithArgs(mockUser.Phone, mockUser.Email, mockUser.Password, mockUser.UserType, mockUser.Verification, mockUser.Status).
					WillReturnRows(sqlmock.NewRows([]string{"id_user"}).AddRow(1))
			},
			expectedId:  1,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQueryFunc()
			ctx := context.Background()

			actualId, actualerr := userRepo.CreateUser(ctx, tt.user)

			assert.Equal(t, tt.expectedErr, actualerr)

			assert.Equal(t, tt.expectedId, actualId)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestUserRepo_FindUserWithNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)

	rows := sqlmock.NewRows([]string{"id_user", "phone", "email", "password", "user_type", "verification", "status"}).
		AddRow(0, "1", "email", "password", "user_type", true, "status")

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
				mock.ExpectQuery("SELECT id_user, phone, email, password, user_type, verification, status from users WHERE phone=\\$1;").
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
				mock.ExpectQuery("SELECT id_user, phone, email, password, user_type, verification, status from users WHERE phone=\\$1;").
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
				mock.ExpectQuery("SELECT id_user, phone, email, password, user_type, verification, status from users WHERE phone=\\$1;").
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

	rows := sqlmock.NewRows([]string{"id_user", "phone", "email", "password", "user_type", "verification", "status"}).
		AddRow(0, "phone", "jon@gmail.com", "password", "user_type", true, "status")

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
				mock.ExpectQuery("SELECT id_user, phone, email, password, user_type, verification, status from users WHERE email=\\$1;").
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
				mock.ExpectQuery("SELECT id_user, phone, email, password, user_type, verification, status from users WHERE email=\\$1;").
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
				mock.ExpectQuery("SELECT id_user, phone, email, password, user_type, verification, status from users WHERE email=\\$1;").
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
