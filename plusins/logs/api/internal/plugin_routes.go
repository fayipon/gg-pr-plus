// gg-pr-plus/plugins/logs/api/internal/plugin_routes.go
package internal

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest"

    "GG-PR/api/internal/svc"
    "GG-PR-PLUS/plugins/logs/api/internal/handler"
)

func RegisterPluginRoutes(server *rest.Server, ctx *svc.ServiceContext) {
    server.AddRoutes(
        rest.WithMiddlewares(
            []rest.Middleware{ctx.Authority},
            []rest.Route{
                { Method: http.MethodPost, Path: "/logs/list", Handler: handler.GetLogsHandler(ctx) },
            }...,
        ),
        rest.WithJwt(ctx.Config.Auth.AccessSecret),
    )
}
