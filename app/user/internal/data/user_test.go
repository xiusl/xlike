package data

import (
    "context"
    "database/sql"
    "errors"
    _ "github.com/go-sql-driver/mysql"
    "github.com/stretchr/testify/require"
    "testing"
    "xlike/app/user/internal/biz"
)


func createTestRepo() *userRepo {
    return &userRepo{
        data: testData,
    }
}

func TestNewUserRepo(t *testing.T) {
    NewUserRepo(testData)
}

func createTestUser(t *testing.T, repo *userRepo, u *biz.User) int64 {
    resID, err := repo.CreateUser(context.Background(), u)
    require.NoError(t, err)
    require.NotZero(t, resID)
    return resID
}

func TestUserRepo_CreateUser(t *testing.T) {
    var u = &biz.User{
        Mobile: "13800000001",
        NickName: "user_0001",
        Avatar: "/default.png",
    }
    repo := createTestRepo()
    createTestUser(t, repo, u)

    _, err := repo.CreateUser(context.Background(), u)
    require.Error(t, err)
}

func TestUserRepo_GetUser(t *testing.T) {
    var u = &biz.User{
        Mobile: "13800000002",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    repo := createTestRepo()

    {   // ok
        resID := createTestUser(t, repo, u)
        res, err := repo.GetUser(context.Background(), resID)
        require.NoError(t, err)
        require.Equal(t, u.Mobile, res.Mobile)
        require.Equal(t, u.NickName, res.NickName)
        require.Equal(t, u.Avatar, res.Avatar)
        require.NotZero(t, res.CreatedAt)
        require.NotZero(t, res.UpdatedAt)
    }


    {   // not found
        res, err := repo.GetUser(context.Background(), 0)
        require.Error(t, err)
        require.Nil(t, res)
        require.True(t, errors.Is(err, sql.ErrNoRows))
    }
}

func TestUserRepo_GetUserByMobile(t *testing.T) {
    var u = &biz.User{
        Mobile: "13800000003",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    repo := createTestRepo()

    {   // ok
        _ = createTestUser(t, repo, u)
        res, err := repo.GetUserByMobile(context.Background(), u.Mobile)
        require.NoError(t, err)
        require.Equal(t, u.Mobile, res.Mobile)
        require.Equal(t, u.NickName, res.NickName)
        require.Equal(t, u.Avatar, res.Avatar)
        require.NotZero(t, res.CreatedAt)
        require.NotZero(t, res.UpdatedAt)
    }


    {   // not found
        res, err := repo.GetUserByMobile(context.Background(), "")
        require.Error(t, err)
        require.Nil(t, res)
        require.True(t, errors.Is(err, sql.ErrNoRows))
    }
}