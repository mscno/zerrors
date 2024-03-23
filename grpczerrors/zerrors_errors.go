package grpczerrors

import (
	"errors"
	"github.com/mscno/zerrors"
	"github.com/samber/oops"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ZNOWToGRPCError(err error) error {
	if err == nil {
		return nil
	}
	code, key, id, ok := ExtractZNOWError(err)
	if !ok {
		return status.Convert(err).Err()
	}
	msg := key
	msg += " (" + id + ")"

	s, err := status.New(code, msg).WithDetails(&ErrorDetail{Id: id, Message: key})
	if err != nil {
		return status.New(code, msg).Err()
	}

	return s.Err()
}

func ExtractZNOWError(err error) (c codes.Code, msg, id string, ok bool) {
	if err == nil {
		return codes.OK, "", "", false
	}
	var domainErr oops.OopsError
	if ok := errors.As(err, &domainErr); !ok {
		return codes.Unknown, err.Error(), "", false
	}
	switch {
	case zerrors.IsAlreadyExists(err):
		return codes.AlreadyExists, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsDeadlineExceeded(err):
		return codes.DeadlineExceeded, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsInternal(err):
		return codes.Internal, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsInvalidArgument(err):
		return codes.InvalidArgument, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsNotFound(err):
		return codes.NotFound, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsPermissionDenied(err):
		return codes.PermissionDenied, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsFailedPrecondition(err):
		return codes.FailedPrecondition, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsUnauthenticated(err):
		return codes.Unauthenticated, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsUnavailable(err):
		return codes.Unavailable, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsUnimplemented(err):
		return codes.Unimplemented, domainErr.Error(), domainErr.Domain(), true
	case zerrors.IsResourceExhausted(err):
		return codes.ResourceExhausted, domainErr.Error(), domainErr.Domain(), true
	default:
		return codes.Unknown, err.Error(), "", false
	}
}
