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

func TestUserRepo_UpdateUser(t *testing.T) {
    var u = &biz.User{
        Mobile: "13800000004",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    repo := createTestRepo()
    resID := createTestUser(t, repo, u)

    {   // ok
        u.ID = resID
        u.NickName = "user_0000"
        res, err := repo.UpdateUser(context.Background(), u)
        require.NoError(t, err)
        require.Equal(t, u.NickName, res.NickName)
        require.Equal(t, u.Mobile, res.Mobile)
        require.Equal(t, u.Avatar, res.Avatar)
    }

    {    // no row affected
        u.ID = 0
        u.NickName = "user_0001"
        res, err := repo.UpdateUser(context.Background(), u)
        require.Error(t, err)
        require.Nil(t, res)
    }
}

func TestUserRepo_DeleteUser(t *testing.T) {
    var u = &biz.User{
        Mobile: "13800000005",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    repo := createTestRepo()
    resID := createTestUser(t, repo, u)
    var err error
    {   // ok
        err = repo.DeleteUser(context.Background(), resID)
        require.NoError(t, err)

        res, err := repo.GetUser(context.Background(), resID)
        require.Error(t, err)
        require.Nil(t, res)
        require.True(t, errors.Is(err, sql.ErrNoRows))
    }

    {   // no row affected
        err = repo.DeleteUser(context.Background(), resID)
        require.Error(t, err)
        require.Equal(t, "no row affected", err.Error())
    }
}

func TestUserRepo_ListUserByIDs(t *testing.T) {
    var u1 = &biz.User{
        Mobile: "13800000006",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    var u2 = &biz.User{
        Mobile: "13800000007",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    repo := createTestRepo()
    ID1 := createTestUser(t, repo, u1)
    ID2 := createTestUser(t, repo, u2)
    { // ok
        us, err := repo.ListUserByIDs(context.Background(), []int64{ID1, ID2})
        require.NoError(t, err)
        require.Equal(t, 2, len(us))
        require.Equal(t, us[0].ID, u1.ID)
        require.Equal(t, us[1].ID, u2.ID)
    }

    { // empty
        us, err := repo.ListUserByIDs(context.Background(), []int64{})
        require.NoError(t, err)
        require.Empty(t, us)
    }

    {  // SQL syntax
        us, err := repo.ListUserByIDs(context.Background(), []int64{ID1, 0})
        require.Error(t, err)
        require.Nil(t, us)
    }

    {  // other
        us, err := repo.ListUserByIDs(context.Background(), []int64{ID1, -1})
        require.NoError(t, err)
        require.Equal(t, 1, len(us))
        require.Equal(t, us[0].ID, u1.ID)
    }
}


func TestUserRepo_ListUserByCursor(t *testing.T) {
    _ = clearTestUserData()
    repo := createTestRepo()
    ID1, ID2, ID3 := createTestUsers(t, repo)

    {
        us, err := repo.ListUserByCursor(context.Background(), ID1, 2)
        require.NoError(t, err)
        require.Len(t, us, 2)
        require.Equal(t, us[0].ID, ID2)
        require.Equal(t, us[1].ID, ID3)
    }

    {
        us, err := repo.ListUserByCursor(context.Background(), ID1, 1)
        require.NoError(t, err)
        require.Len(t, us, 1)
        require.Equal(t, us[0].ID, ID2)
    }

    {
        us, err := repo.ListUserByCursor(context.Background(), ID3, 1)
        require.NoError(t, err)
        require.Len(t, us, 0)
    }

    {
       us, err := repo.ListUserByCursor(context.Background(), 0, 1)
       require.NoError(t, err)
       require.Len(t, us, 1)
       require.Equal(t, us[0].ID, ID1)
    }

    {
       us, err := repo.ListUserByCursor(context.Background(), 0, -1)
       require.NoError(t, err)
       require.Len(t, us, 3)
       require.Equal(t, us[0].ID, ID1)
    }
}

func createTestUsers(t *testing.T, repo *userRepo) (int64, int64, int64) {
    var u1 = &biz.User{
        Mobile: "13800000008",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    var u2 = &biz.User{
        Mobile: "13800000009",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    var u3 = &biz.User{
        Mobile: "13800000010",
        NickName: "user_0002",
        Avatar: "/default.png",
    }
    _ = clearTestUserData()
    ID1 := createTestUser(t, repo, u1)
    ID2 := createTestUser(t, repo, u2)
    ID3 := createTestUser(t, repo, u3)
    return ID1, ID2, ID3
}

func TestUserRepo_ListUserByCursorDesc(t *testing.T) {
    _ = clearTestUserData()
    repo := createTestRepo()
    ID1, ID2, ID3 := createTestUsers(t, repo)

    {
        us, err := repo.ListUserByCursorDesc(context.Background(), 0, 3)
        require.NoError(t, err)
        require.Len(t, us, 3)
        require.Equal(t, us[0].ID, ID3)
    }

    {
        us, err := repo.ListUserByCursorDesc(context.Background(), ID3, 2)
        require.NoError(t, err)
        require.Len(t, us, 2)
        require.Equal(t, us[0].ID, ID2)
        require.Equal(t, us[1].ID, ID1)
    }

    {
        us, err := repo.ListUserByCursorDesc(context.Background(), ID1, 2)
        require.NoError(t, err)
        require.Len(t, us, 0)
    }
}
