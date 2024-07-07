// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/atomicals-go/atomicals-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/checkTx",
				Handler: checkTxHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/getAssetByLocationID",
				Handler: getAssetByLocationIDHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/getAssetByUserPk",
				Handler: getAssetByUserPkHandler(serverCtx),
			},
		},
	)
}