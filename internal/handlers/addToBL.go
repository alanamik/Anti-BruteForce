package handlers

import (
	"OTUS_hws/Anti-BruteForce/internal/gen/restapi/operations"
	"OTUS_hws/Anti-BruteForce/internal/redisdb"

	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) AddToBL(params operations.BlacklistPutParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()
	err := h.redisServer.AddToList(ctx, params.HTTPRequest.Host, redisdb.Blacklist)
	if err != nil {
		return operations.NewBlacklistDeleteInternalServerError().WithPayload(&operations.BlacklistDeleteInternalServerErrorBody{
			Ok: err.Error(),
		})
	}

	return operations.NewBlacklistDeleteOK().WithPayload(&operations.BlacklistDeleteOKBody{
		Ok: "The IP has been successfully added to the list",
	})
}
