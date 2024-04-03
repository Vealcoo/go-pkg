package httphelper

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
)

// ErrorMessage general error message
type ErrorMessage struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Args         interface{} `json:"args,omitempty"`
	Status       string      `json:"status"`
	MessageLangx string      `json:"messageLangx,omitempty"`
}

// CreateErrorMessage create error message
func CreateErrorMessage(c codes.Code, mesg string, args interface{}) *ErrorMessage {
	return &ErrorMessage{
		Code:    HTTPStatusFromCode(c),
		Message: mesg,
		Status:  CodeName[uint32(c)],
		Args:    args,
	}
}

// CodeName grpc code an name mapping
var CodeName = map[uint32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	16: "UNAUTHENTICATED",
	8:  "RESOURCE_EXHAUSTED",
	9:  "FAILED_PRECONDITION",
	10: "ABORTED",
	11: "OUT_OF_RANGE",
	12: "UNIMPLEMENTED",
	13: "INTERNAL",
	14: "UNAVAILABLE",
	15: "DATA_LOSS",
}

// HTTPStatusFromCode trans grpc code to http code
func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}

	log.Info().Str("Message", "Unknown gRPC error code").Uint32("code", uint32(code))
	return http.StatusInternalServerError
}
