// gg-pr-plus/plugins/logs/api/internal/types/types.go
package types

type LogsListReq struct {
    Page     int `json:"page"`
    PageSize int `json:"pageSize"`
}

type LogRecord struct {
    Id         int64  `json:"id"`
    UserId     int64  `json:"user_id"`
    Action     string `json:"action"`
    Detail     string `json:"detail"`
    Ip         string `json:"ip"`
    CreateTime string `json:"create_time"`
}

type LogsListResp struct {
    Total int64       `json:"total"`
    Data  []LogRecord `json:"data"`
}
