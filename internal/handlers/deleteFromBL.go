package handlers

import (
	"OTUS_hws/Anti-BruteForce/internal/gen/restapi/operations"
	"OTUS_hws/Anti-BruteForce/internal/redisdb"

	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) DeleteFromBL(params operations.BlacklistDeleteParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	err := h.redisServer.DeleteFromList(ctx, params.HTTPRequest.Host, redisdb.Blacklist)
	if err != nil {
		return operations.NewBlacklistDeleteInternalServerError().WithPayload(&operations.BlacklistDeleteInternalServerErrorBody{
			Ok: err.Error(),
		})
	}

	return operations.NewBlacklistDeleteOK().WithPayload(&operations.BlacklistDeleteOKBody{
		Ok: "The IP has been successfully deleted from the list",
	})
}
