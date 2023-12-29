package swagger

import (
	httputil "github.com/vuduongtp/go-core/pkg/util/http"
	_ "github.com/vuduongtp/go-core/pkg/util/swagger" // Swagger stuffs
)

// ListRequest holds data of listing request from react-admin
type ListRequest struct {
	httputil.ListRequest
}
