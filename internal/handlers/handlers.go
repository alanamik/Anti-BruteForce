package handlers

import (
	"OTUS_hws/Anti-BruteForce/internal/gen/restapi/operations"
	"OTUS_hws/Anti-BruteForce/internal/redisdb"
	"time"
)

const (
	requestsPerMinutes = 10
)

type Bucket struct {
	Blocked  bool
	N        int64
	Requests []time.Time
}

type Handler struct {
	redisServer *redisdb.RedisClient
	clients     map[string]Bucket
}

func NewHandler(redisS *redisdb.RedisClient) *Handler {
	return &Handler{
		redisServer: redisS,
		clients:     make(map[string]Bucket),
	}
}

func (h *Handler) Register(api *operations.AntiBrutForceAPI) {
	api.AuthCheckHandler = operations.AuthCheckHandlerFunc(h.AuthCheck)
	api.BlacklistDeleteHandler = operations.BlacklistDeleteHandlerFunc(h.DeleteFromBL)
	api.BlacklistPutHandler = operations.BlacklistPutHandlerFunc(h.AddToBL)
}
