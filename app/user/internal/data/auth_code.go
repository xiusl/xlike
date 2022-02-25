package data

import (
    "context"
    "time"
    "xlike/app/user/internal/biz"
)

type authCode struct {
    id        int64
    mobile    string
    code      string
    bizCode   string
    status    int
    createdAt time.Time
    expiredAt time.Time
}

type authCodeRepo struct {
    data *Data
}

func NewAuthCodeRepo(data *Data) biz.CodeRepo {
    return &authCodeRepo{
        data: data,
    }
}

var (
    createAuthCodeSql = "insert into auth_code (key, code, biz_code, created_at, expired_at) values (?, ?, ?, ?, ?)"
    getAuthCodeSql    = "select id, `key`, code, biz_code, status, created_at, expired_at from auth_code where `key`=? and status=0 order by created_at desc limit 1"
    updateAuthCodeSql = "update auth_code set status=1 where key=?"
)

func toBizAuthCode(c *authCode) *biz.AuthCode {
    return &biz.AuthCode{
        ID:        c.id,
        Mobile:    c.mobile,
        Code:      c.code,
        BizCode:   c.bizCode,
        Status:    c.status,
        CreatedAt: c.createdAt,
        ExpiredAt: c.expiredAt,
    }
}

func (r *authCodeRepo) GetAuthCodeByMobile(ctx context.Context, mobile string) (*biz.AuthCode, error) {
    row := r.data.db.QueryRowContext(ctx, getAuthCodeSql, mobile)
    var c authCode
    err := row.Scan(&c.id, &c.mobile, &c.code, &c.bizCode, &c.status, &c.createdAt, &c.expiredAt)
    if err != nil {
        return nil, err
    }
    return toBizAuthCode(&c), nil
}

func (r *authCodeRepo) CreateAuthCode(ctx context.Context, c *biz.AuthCode) (int64, error) {
    res, err := r.data.db.ExecContext(ctx, createAuthCodeSql, c.Mobile, c.Code, c.BizCode, c.CreatedAt, c.ExpiredAt)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}
