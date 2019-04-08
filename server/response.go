package server

const (
	SUCCESS string = "success"
	ERROR   string = "error"
)

type response struct {
	Status string `json:"status"`
}

type successResponse struct {
	response
	ShortenedUrl string `json:"shortened_url"`
	OriginalUrl  string `json:"original_url"`
}

type errorResponse struct {
	response
	Message string `json:"message"`
}

func newSuccessResponse(shortened string, original string) *successResponse {
	return &successResponse{response: response{SUCCESS}, ShortenedUrl: shortened, OriginalUrl: original}
}

func newErrorResponse(message string) *errorResponse {
	return &errorResponse{response: response{ERROR}, Message: message}
}
