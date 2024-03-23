package grpczerrors

import (
	"errors"
	"github.com/mscno/zerrors"

	"testing"

	"google.golang.org/grpc/codes"
)

func TestCaosToGRPCError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"no error",
			args{},
			false,
		},
		{
			"unknown error",
			args{errors.New("unknown")},
			true,
		},
		{
			"caos error",
			args{zerrors.ToInternal(nil, "", "message")},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ZNOWToGRPCError(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("ZNOWToGRPCError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Extract(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantC   codes.Code
		wantMsg string
		wantID  string
		wantOk  bool
	}{
		{
			"already exists",
			args{zerrors.ToAlreadyExists(nil, "id", "already exists")},
			codes.AlreadyExists,
			"already exists",
			"id",
			true,
		},
		{
			"deadline exceeded",
			args{zerrors.ToDeadlineExceeded(nil, "id", "deadline exceeded")},
			codes.DeadlineExceeded,
			"deadline exceeded",
			"id",
			true,
		},
		{
			"internal error",
			args{zerrors.ToInternal(nil, "id", "internal error")},
			codes.Internal,
			"internal error",
			"id",
			true,
		},
		{
			"invalid argument",
			args{zerrors.ToInvalidArgument(nil, "id", "invalid argument")},
			codes.InvalidArgument,
			"invalid argument",
			"id",
			true,
		},
		{
			"not found",
			args{zerrors.ToNotFound(nil, "id", "not found")},
			codes.NotFound,
			"not found",
			"id",
			true,
		},
		{
			"permission denied",
			args{zerrors.ToPermissionDenied(nil, "id", "permission denied")},
			codes.PermissionDenied,
			"permission denied",
			"id",
			true,
		},
		{
			"precondition failed",
			args{zerrors.ToFailedPrecondition(nil, "id", "precondition failed")},
			codes.FailedPrecondition,
			"precondition failed",
			"id",
			true,
		},
		{
			"unauthenticated",
			args{zerrors.ToUnauthenticated(nil, "id", "unauthenticated")},
			codes.Unauthenticated,
			"unauthenticated",
			"id",
			true,
		},
		{
			"unavailable",
			args{zerrors.ToUnavailable(nil, "id", "unavailable")},
			codes.Unavailable,
			"unavailable",
			"id",
			true,
		},
		{
			"unimplemented",
			args{zerrors.ToUnimplemented(nil, "id", "unimplemented")},
			codes.Unimplemented,
			"unimplemented",
			"id",
			true,
		},
		{
			"exhausted",
			args{zerrors.ToResourceExhausted(nil, "id", "exhausted")},
			codes.ResourceExhausted,
			"exhausted",
			"id",
			true,
		},
		{
			"unknown",
			args{errors.New("unknown")},
			codes.Unknown,
			"unknown",
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, gotMsg, gotID, gotOk := ExtractZNOWError(tt.args.err)
			if gotC != tt.wantC {
				t.Errorf("extract() gotC = %v, want %v", gotC, tt.wantC)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("extract() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
			if gotID != tt.wantID {
				t.Errorf("extract() gotID = %v, want %v", gotID, tt.wantID)
			}
			if gotOk != tt.wantOk {
				t.Errorf("extract() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
