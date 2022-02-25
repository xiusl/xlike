package biz

import (
    "context"
    "database/sql"
    "errors"
)



func (uc *UseCase) Register(ctx context.Context, mobile, code string) (*User, string, error) {
    return nil, "", nil
}

func (uc *UseCase) Auth(ctx context.Context, mobile, code string) (*User, string, error) {
    _, err := uc.GetUserByMobile(ctx, mobile)
    if err != nil && !errors.Is(err, sql.ErrNoRows) {
        return uc.Login(ctx, mobile, code)
    }
    return uc.Register(ctx, mobile, code)
}

func (uc *UseCase) Login(ctx context.Context, mobile, code string) (*User, string, error) {
    return nil, "", nil
}

func (uc *UseCase) VerifyAuthCode(ctx context.Context, mobile, code string) error {
    return nil
}