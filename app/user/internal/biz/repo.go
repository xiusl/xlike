package biz

import "context"

type UserRepo interface {
    CreateUser(ctx context.Context, u *User) (int64, error)
    GetUser(ctx context.Context, ID int64) (*User, error)
    GetUserByMobile(ctx context.Context, mobile string) (*User, error)

}

type CodeRepo interface {
    GetAuthCodeByMobile(ctx context.Context, mobile string) (*AuthCode, error)
}
