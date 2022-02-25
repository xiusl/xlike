package biz

import (
    "context"
    "time"
)

type User struct {
    ID        int64
    Mobile    string
    NickName  string
    Avatar    string
    State     int
    Memo      string
    LastSeen  time.Time
    Version   int
    Del       int
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (uc *UseCase) CreateUser(ctx context.Context, u *User) (*User, error) {
    ID, err := uc.repo.CreateUser(ctx, u)
    if err != nil {
        return nil, err
    }
    return uc.repo.GetUser(ctx, ID)
}

func (uc *UseCase) GetUser(ctx context.Context, ID int64) (*User, error) {
    return uc.repo.GetUser(ctx, ID)
}

func (uc *UseCase) GetUserByMobile(ctx context.Context, mobile string) (*User, error) {
    return uc.repo.GetUserByMobile(ctx, mobile)
}