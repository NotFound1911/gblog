package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/NotFound1911/morm"
	"github.com/lib/pq"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = sql.ErrNoRows
)

//go:generate mockgen -source=./user.go -package=daomocks -destination=./mocks/user.mock.go UserDAO
type UserDAO interface {
	Insert(ctx context.Context, u User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	UpdateById(ctx context.Context, entity User) error
	FindById(ctx context.Context, uid int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
}

type OrmUser struct {
	db *morm.DB
}

var _ UserDAO = &OrmUser{}

func (o OrmUser) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	res := morm.NewInserter[User](o.db).Values(&u).Exec(ctx)
	err := res.Err()
	if pe, ok := err.(*pq.Error); ok {
		const UniqueViolation pq.ErrorCode = "23505"
		if pe.Code == UniqueViolation {
			// 用户冲突，邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return res.Err()
}

func (o OrmUser) FindByEmail(ctx context.Context, email string) (User, error) {
	res, err := morm.NewSelector[User](o.db).Where(morm.C("Email").EQ(email)).Get(ctx)
	return *res, err
}

func (o OrmUser) UpdateById(ctx context.Context, entity User) error {
	res := morm.NewUpdater[User](o.db).Update(&entity).
		Set(morm.C("Email"), morm.C("Password"), morm.C("Nickname"), morm.C("Birthday"),
			morm.C("AboutMe"), morm.C("Phone"), morm.Assign("Utime", time.Now().UnixMilli())).
		Where(morm.C("Id").EQ(entity.Id)).Exec(ctx)
	return res.Err()
}

func (o OrmUser) FindById(ctx context.Context, uid int64) (User, error) {
	res, err := morm.NewSelector[User](o.db).Where(morm.C("Id").EQ(uid)).Get(ctx)
	return *res, err
}

func (o OrmUser) FindByPhone(ctx context.Context, phone string) (User, error) {
	res, err := morm.NewSelector[User](o.db).Where(morm.C("Phone").EQ(phone)).Get(ctx)
	return *res, err
}

func NewOrmUser(db *morm.DB) UserDAO {
	return &OrmUser{
		db: db,
	}
}

type User struct {
	Id int64 `morm:"column=id"`

	Email    sql.NullString `morm:"column=email"`
	Password string         `morm:"column=password"`

	Nickname string `morm:"column=nickname"`
	// YYYY-MM-DD
	Birthday int64  `morm:"column=birthday"`
	AboutMe  string `morm:"column=about_me"`

	// 代表这是一个可以为 NULL 的列
	Phone sql.NullString `morm:"column=phone"`

	// 时区，UTC 0 的毫秒数
	// 创建时间
	Ctime int64 `morm:"column=c_time"`
	// 更新时间
	Utime int64 `morm:"column=u_time"`
}
