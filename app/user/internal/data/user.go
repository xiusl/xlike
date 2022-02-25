package data

import (
    "context"
    "database/sql"
    "fmt"
    "time"
    "xlike/app/user/internal/biz"
    "xlike/pkg/utils"
)

var (
    createUserSql = "insert into user (id, mobile, nick_name, avatar) values (?, ?, ?, ?)"
    getUserSql    = "select id, mobile, nick_name, avatar, state, memo, last_seen, version, del, created_at, updated_at from user where id=? and del=0"
    getUserByMobileSql    = "select id, mobile, nick_name, avatar, state, memo, last_seen, version, del, created_at, updated_at from user where mobile=? and del=0"
    updateUserSql = "update user set mobile=?, nick_name=?, avatar=?, state=?, memo=? where id=?"
    deleteUserSql = "update user set del=1 where id=?"
    listUserByCursorSql = "select id, mobile, nick_name, avatar, state, memo, last_seen, version, del, created_at, updated_at from user where id > ? limit ?"
    listUserByCursorDescSql = "select id, mobile, nick_name, avatar, state, memo, last_seen, version, del, created_at, updated_at from user where id < ? limit ?"
    listUserByIDsSql = "select id, mobile, nick_name, avatar, state, memo, last_seen, version, del, created_at, updated_at from user where id in (%s)"
)

type user struct {
    id        int64
    mobile    string
    nickName  string
    avatar    sql.NullString
    state     int
    memo      sql.NullString
    lastSeen  sql.NullTime
    version   int
    del       int
    createdAt time.Time
    updatedAt time.Time
}

func toBizUser(u *user) *biz.User {
    return &biz.User{
        ID:        u.id,
        Mobile:    u.mobile,
        NickName:  u.nickName,
        Avatar:    u.avatar.String,
        State:     u.state,
        Memo:      u.memo.String,
        LastSeen:  u.lastSeen.Time,
        Version:   u.version,
        Del:       u.del,
        CreatedAt: u.createdAt,
        UpdatedAt: u.updatedAt,
    }
}

type userRepo struct {
    data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
    return &userRepo{
        data: data,
    }
}

// CreateUser is .
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (int64, error) {
    if u.ID == 0 {
        u.ID = int64(r.data.node.Generate())
    }
    _, err := r.data.db.ExecContext(ctx, createUserSql, u.ID, u.Mobile, u.NickName, u.Avatar)
    if err != nil {
        return 0, err
    }
    return u.ID, nil
}

// GetUser is .
func (r *userRepo) GetUser(ctx context.Context, ID int64) (*biz.User, error) {
    row := r.data.db.QueryRowContext(ctx, getUserSql, ID)
    u, err := scanUser(row)
    if err != nil {
        return nil, err
    }
    return toBizUser(u), nil
}


// GetUserByMobile is .
func (r *userRepo) GetUserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
    row := r.data.db.QueryRowContext(ctx, getUserByMobileSql, mobile)
    u, err := scanUser(row)
    if err != nil {
        return nil, err
    }
    return toBizUser(u), nil
}

// UpdateUser is .
func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
    res, err := r.data.db.ExecContext(ctx, updateUserSql, u.Mobile, u.NickName, u.Avatar, u.State, u.Memo, u.ID)
    if err != nil {
        return nil, err
    }
    _, err = res.RowsAffected()
    if err != nil {
        return nil, err
    }
    return u, nil
}

// DeleteUser is logic delete
func (r *userRepo) DeleteUser(ctx context.Context, ID int64) error {
    res, err := r.data.db.ExecContext(ctx, deleteUserSql, ID)
    if err != nil {
        return err
    }
    _, err = res.RowsAffected()
    if err != nil {
        return err
    }
    return nil
}

// ListUserByIDs is .
func (r *userRepo) ListUserByIDs(ctx context.Context, IDs []int64) ([]*biz.User, error) {
    sqlStr := fmt.Sprintf(listUserByIDsSql, utils.IntsToStrs(IDs))
    rows, err := r.data.db.QueryContext(ctx, sqlStr)
    if err != nil {
        return nil, err
    }
    return scanUsers(rows)
}

// ListUserByCursor is
func (r *userRepo) ListUserByCursor(ctx context.Context, cursor int64, count int) ([]*biz.User, error) {
    rows, err := r.data.db.QueryContext(ctx, listUserByCursorSql, cursor, count)
    if err != nil {
        return nil, err
    }
    return scanUsers(rows)
}

// ListUserByCursorDesc is
func (r *userRepo) ListUserByCursorDesc(ctx context.Context, cursor int64, count int) ([]*biz.User, error) {
    rows, err := r.data.db.QueryContext(ctx, listUserByCursorDescSql, cursor, count)
    if err != nil {
        return nil, err
    }
    return scanUsers(rows)
}
func scanUser(row *sql.Row) (*user, error) {
    var u user
    err := row.Scan(
        &u.id,
        &u.mobile,
        &u.nickName,
        &u.avatar,
        &u.state,
        &u.memo,
        &u.lastSeen,
        &u.version,
        &u.del,
        &u.createdAt,
        &u.updatedAt,
    )
    if err != nil {
        return nil, err
    }
    return &u, nil
}
func scanUsers(rows *sql.Rows) ([]*biz.User, error) {
    rs := make([]*biz.User, 0)
    for rows.Next() {
        var u user
        err := rows.Scan(
            &u.id,
            &u.mobile,
            &u.nickName,
            &u.avatar,
            &u.state,
            &u.memo,
            &u.lastSeen,
            &u.version,
            &u.del,
            &u.createdAt,
            &u.updatedAt,
            )
        if err != nil {
            return nil, err
        }
        rs = append(rs, toBizUser(&u))
    }
    return rs, nil
}