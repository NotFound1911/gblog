package cache

import (
	"context"
	"fmt"
	"github.com/NotFound1911/gblog/internal/user/domain"
	mocks "github.com/NotFound1911/gblog/internal/user/repository/cache/mocks"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisUserCache_Get(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) redis.Cmdable
		ctx  context.Context
		uid  int64
		user domain.User

		wantErr error
	}{
		{
			name: "no key",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				res := redis.NewStringResult("", redis.Nil)
				cmd.EXPECT().Get(context.Background(), testGenKey(123)).Return(res)
				return cmd
			},
			wantErr: ErrKeyNotExist,
			uid:     123,
			ctx:     context.Background(),
		},
		{
			name: "normal data",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				res := redis.NewStringResult(`{
	"id": 123,
	"email": "456@qq.com",
	"nickname": "test",
	"password": "!@#",
	"phone": "1111111",
	"about_me": "123",
	"ctime": "2006-01-02T15:04:05+08:00",
	"birthday": "2006-01-02T15:04:05+08:00"
}`, nil)
				cmd.EXPECT().Get(context.Background(), testGenKey(123)).Return(res)
				return cmd
			},
			uid: 123,
			ctx: context.Background(),
			user: func() domain.User {
				u := domain.User{
					Id:       123,
					Email:    "456@qq.com",
					Nickname: "test",
					Password: "!@#",
					Phone:    "1111111",
					AboutMe:  "123",
				}
				ctime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+08:00")
				birthday, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+08:00")
				u.Ctime = ctime
				u.Birthday = birthday
				return u
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cache := NewUserCache(tc.mock(ctrl))
			u, err := cache.Get(tc.ctx, tc.uid)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.user, u)
		})
	}
}
func testGenKey(uid int64) string {
	return fmt.Sprintf("user:key:%d", uid)
}
