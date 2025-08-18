package error_usecase

var (
	InternalServerError = "internal server error"
	RouteNotFound       = "route not found"
	InvalidRequestBody  = "invalid request body"
	InvalidParams       = "invalid params"
	InvalidQueryParams  = "invalid query params"
	InvalidFormData     = "invalid form data"
	Unauthorized        = "unauthorized"
	Forbidden           = "forbidden"
	ServiceUnavailable  = "service unavailable"
	ProcessTimeout      = "process timeout"
)

type ErrorsResp struct {
	Messages []string `json:"messages"`
}

type ErrorResp struct {
	Message string `json:"message"`
}
