package swaggerutil

import (
	"github.com/vuduongtp/go-core/pkg/server"
)

// SwaggOKResp success empty response
type SwaggOKResp struct{}

// SwaggErrResp error empty response
type SwaggErrResp struct{}

// SwaggErrDetailsResp model error response
type SwaggErrDetailsResp struct {
	server.ErrorResponse
}
