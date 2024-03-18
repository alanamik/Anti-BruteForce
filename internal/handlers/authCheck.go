package handlers

import (
	"OTUS_hws/Anti-BruteForce/internal/gen/restapi/operations"
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

func authCheckError() middleware.Responder {
	return operations.NewAuthCheckInternalServerError().WithPayload(&operations.AuthCheckInternalServerErrorBody{Ok: false})
}

func (h *Handler) AuthCheck(params operations.AuthCheckParams) middleware.Responder {
	userIp := params.HTTPRequest.Host
	v, ok := h.clients[userIp]
	if !ok {
		client := Bucket{
			Blocked:  false,
			N:        1,
			Requests: make([]time.Time, requestsPerMinutes),
		}
		client.Requests[0] = time.Now()

		h.clients[userIp] = client
		fmt.Println(v.Blocked)
	} else {
		newC := Bucket{
			Blocked:  v.Blocked,
			N:        v.N,
			Requests: v.Requests,
		}
		fmt.Println(newC.Blocked, newC.N, newC.Requests)
		if newC.Blocked {
			return authCheckError()
		}

		if newC.N >= requestsPerMinutes {
			newC.N = 0
		}
		next := newC.N + 1

		if next >= requestsPerMinutes {
			next = 0
		}

		if time.Since(newC.Requests[next]) < time.Minute {
			newC.Blocked = true
			fmt.Println("ASD", newC.Blocked)
			return authCheckError()
		}

		newC.Requests[newC.N] = time.Now()

		newC.N++

		h.clients[userIp] = newC
	}

	return operations.NewAuthCheckOK().WithPayload(&operations.AuthCheckOKBody{Ok: true})
}
