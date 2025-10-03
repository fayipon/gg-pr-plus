// gg-pr-plus/plugins/logs/api/internal/handler/get_logs_handler.go
package handler

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
    "GG-PR-PLUS/plugins/logs/api/internal/logic"
    "GG-PR-PLUS/plugins/logs/api/internal/types"
    "GG-PR/api/internal/svc"
)

func GetLogsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.LogsListReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }
        l := logic.NewGetLogsLogic(r.Context(), ctx)
        resp, err := l.GetLogs(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }
        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
