package grpczerrors

import (
	"errors"
	"github.com/mscno/zerrors"

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
	zitadelErr := new(zerrors.Zerror)
	if ok := errors.As(err, &zitadelErr); !ok {
		return codes.Unknown, err.Error(), "", false
	}
	switch {
	case zerrors.IsAlreadyExists(err):
		return codes.AlreadyExists, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsDeadlineExceeded(err):
		return codes.DeadlineExceeded, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsInternal(err):
		return codes.Internal, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsInvalidArgument(err):
		return codes.InvalidArgument, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsNotFound(err):
		return codes.NotFound, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsPermissionDenied(err):
		return codes.PermissionDenied, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsPreconditionFailed(err):
		return codes.FailedPrecondition, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsUnauthenticated(err):
		return codes.Unauthenticated, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsUnavailable(err):
		return codes.Unavailable, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsUnimplemented(err):
		return codes.Unimplemented, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	case zerrors.IsResourceExhausted(err):
		return codes.ResourceExhausted, zitadelErr.GetMessage(), zitadelErr.GetID(), true
	default:
		return codes.Unknown, err.Error(), "", false
	}
}
