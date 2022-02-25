package data

import (
    "context"
    "database/sql"
    "github.com/bwmarrin/snowflake"
    "github.com/go-kratos/kratos/v2/log"
    _ "github.com/go-sql-driver/mysql"
    "os"
    "testing"
)

var (
    testData *Data
    driver = "mysql"
    source = "root:openIM@tcp(192.168.0.23:13306)/xlike_user?parseTime=True&loc=Asia%2FShanghai"
)

func TestMain(m *testing.M) {
    db, err := sql.Open(driver, source)
    if err != nil {
        log.Fatal(err)
        return
    }
    node, err := snowflake.NewNode(1)
    if err != nil {
        log.Fatal(err)
        return
    }
    testData = &Data{
        node: node,
        db: db,
        log: log.NewHelper(log.DefaultLogger),
    }

    _, err = testData.db.ExecContext(context.Background(), "delete from user")
    if err != nil {
        log.Fatal(err)
        return
    }


    os.Exit(m.Run())
}
