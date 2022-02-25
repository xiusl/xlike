package data

import (
	"database/sql"
    "github.com/bwmarrin/snowflake"
    "github.com/go-kratos/kratos/v2/log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/google/wire"
    "xlike/app/user/internal/conf"
)

// ProviderSet is data providers
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data is .
type Data struct {
    node *snowflake.Node
    db *sql.DB
    log *log.Helper
}

// NewData is .
func NewData(c *conf.Data, logger log.Logger) *Data {
    node, err := snowflake.NewNode(1)
    if err != nil {
        panic(err)
    }
    db, err := sql.Open(c.Database.Driver, c.Database.Source)
    if err != nil {
        panic(err)
    }
    return &Data{
        node: node,
        db: db,
        log: log.NewHelper(log.With(logger, "module", "user/data")),
    }
}

func (data *Data) Close() {
    _ = data.db.Close()
}