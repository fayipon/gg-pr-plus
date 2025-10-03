// gg-pr-plus/plugins/logs/api/internal/logic/get_logs_logic.go
package logic

import (
    "context"
    "database/sql"

    "GG-PR/api/internal/svc"
    "GG-PR-PLUS/plugins/logs/api/internal/types"
)

type GetLogsLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogsLogic {
    return &GetLogsLogic{ctx: ctx, svcCtx: svcCtx}
}

func (l *GetLogsLogic) GetLogs(req *types.LogsListReq) (*types.LogsListResp, error) {
    if req.Page < 1 {
        req.Page = 1
    }
    if req.PageSize < 1 {
        req.PageSize = 20
    }
    offset := (req.Page - 1) * req.PageSize
    db := l.svcCtx.DB

    rows, err := db.QueryContext(l.ctx,
        `SELECT id, user_id, action, detail, ip, create_time
         FROM sys_log
         ORDER BY id DESC
         LIMIT ? OFFSET ?`, req.PageSize, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []types.LogRecord
    for rows.Next() {
        var rec types.LogRecord
        if err := rows.Scan(&rec.Id, &rec.UserId, &rec.Action, &rec.Detail, &rec.Ip, &rec.CreateTime); err != nil {
            return nil, err
        }
        list = append(list, rec)
    }

    var total sql.NullInt64
    if err := db.QueryRowContext(l.ctx, `SELECT COUNT(*) FROM sys_log`).Scan(&total); err != nil {
        return nil, err
    }

    return &types.LogsListResp{
        Total: total.Int64,
        Data:  list,
    }, nil
}
