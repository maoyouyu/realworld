package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

type UserUsercase struct {
	user    UserRepo
	profile ProfileRepo

	log *log.Helper
}

func NewUsercase(user UserRepo, profile ProfileRepo, logger log.Logger) *UserUsercase {
	return &UserUsercase{user: user, profile: profile, log: log.NewHelper(logger)}
}

func (uc *UserUsercase) Register(ctx context.Context,user *User) error {
	if err := uc.user.CreateUser(ctx,user);err != nil{
		panic(err)
	}
	return nil
}
