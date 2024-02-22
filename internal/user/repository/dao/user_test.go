package dao

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NotFound1911/morm"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrmUser_Insert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(t *testing.T) *sql.DB
		ctx  context.Context
		user User

		wantErr error
	}{
		{
			name: "插入成功",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mockRes := sqlmock.NewResult(10, 1)
				mock.ExpectExec("INSERT INTO .*").WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "zero", sqlmock.AnyArg(),
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
				).WillReturnResult(mockRes)
				return db
			},
			ctx: context.Background(),
			user: User{
				Nickname: "zero",
			},
		},
		{
			name: "邮箱冲突",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mock.ExpectExec("INSERT INTO .*").WillReturnError(&pq.Error{Code: "23505"})
				return db
			},
			ctx: context.Background(),
			user: User{
				Nickname: "Tom",
			},
			wantErr: ErrDuplicateEmail,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sqlDB := tc.mock(t)
			db, err := morm.OpenDB(sqlDB)
			assert.NoError(t, err)
			dao := NewOrmUser(db)
			err = dao.Insert(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
