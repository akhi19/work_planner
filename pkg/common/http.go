package common

import (
	"context"
	"encoding/json"
	"net/http"
)

type Code int

const (
	serverErrorCode  Code = 500
	paramMissingCode Code = 400
)

const (
	InvalidRequestMsg = "Invalid Request"
)

type Entity struct {
	HttpStatusCode int
	Message        string
	Code           Code
}

func (ue *Entity) Error() string {
	return ue.Message
}

func ValidationError(msg string) error {
	return &Entity{
		HttpStatusCode: http.StatusBadRequest,
		Message:        msg,
		Code:           paramMissingCode,
	}
}

func InternalServerError() error {
	return &Entity{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "Something Went Wrong",
		Code:           serverErrorCode,
	}
}

func SendHttpError(ctx context.Context, w http.ResponseWriter, err error) {
	httpStatusCode := http.StatusInternalServerError
	code := serverErrorCode
	message := err.Error()

	if errVal, ok := err.(*Entity); ok {
		httpStatusCode = errVal.HttpStatusCode
		message = errVal.Message
		code = errVal.Code
	}

	resp := map[string]interface{}{
		"message": message,
		"code":    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(resp)
}

func SendHttpResponse(
	responseWriter http.ResponseWriter,
	statusCode int,
	response interface{},
) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(response)
}

type emptyHttpResponse struct{}

func SendEmptyHttpResponse(
	responseWriter http.ResponseWriter,
	statusCode int,
) {
	responseData := emptyHttpResponse{}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(responseData)
}

func RecoverHandler(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Panic(err)
				serverErr := InternalServerError()
				SendHttpError(r.Context(), w, serverErr)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func AuthenticateHandler(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Add authentication logic
		next.ServeHTTP(w, r)
	}
}

func HttpRequestHandler(
	next http.HandlerFunc,
) http.HandlerFunc {
	next = AuthenticateHandler(next)
	next = RecoverHandler(next)
	return next
}
