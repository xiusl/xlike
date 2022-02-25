package biz

import "time"

type AuthCode struct {
    ID        int64
    Mobile    string
    Code      string
    BizCode   string
    Status    int
    CreatedAt time.Time
    ExpiredAt time.Time
}
