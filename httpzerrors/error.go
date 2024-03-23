package http

import (
	"errors"
	"github.com/mscno/zerrors"
	"github.com/samber/oops"
	"net/http"
)

func ZitadelErrorToHTTPStatusCode(err error) (statusCode int, ok bool) {
	if err == nil {
		return http.StatusOK, true
	}
	//nolint:errorlint
	var domainErr oops.OopsError
	if errors.As(err, &domainErr) {
		switch domainErr.Code() {
		case zerrors.ErrAlreadyExists:
			return http.StatusConflict, true
		case zerrors.ErrDeadlineExceeded:
			return http.StatusGatewayTimeout, true
		case zerrors.ErrInternal:
			return http.StatusInternalServerError, true
		case zerrors.ErrInvalidArgument:
			return http.StatusBadRequest, true
		case zerrors.ErrNotFound:
			return http.StatusNotFound, true
		case zerrors.ErrPermissionDenied:
			return http.StatusForbidden, true
		case zerrors.ErrFailedPrecondition:
			return http.StatusBadRequest, true
		case zerrors.ErrUnauthenticated:
			return http.StatusUnauthorized, true
		case zerrors.ErrUnavailable:
			return http.StatusServiceUnavailable, true
		case zerrors.ErrUnimplemented:
			return http.StatusNotImplemented, true
		case zerrors.ErrResourceExhausted:
			return http.StatusTooManyRequests, true
		default:
			return http.StatusInternalServerError, false
		}
	}
	return http.StatusInternalServerError, false
}
