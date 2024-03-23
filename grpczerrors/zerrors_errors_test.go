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
			args{zerrors.Internal("message")},
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
			args{zerrors.AlreadyExists("already exists")},
			codes.AlreadyExists,
			"already exists",
			"",
			true,
		},
		{
			"deadline exceeded",
			args{zerrors.DeadlineExceeded("deadline exceeded")},
			codes.DeadlineExceeded,
			"deadline exceeded",
			"",
			true,
		},
		{
			"internal error",
			args{zerrors.Internal("internal error")},
			codes.Internal,
			"internal error",
			"",
			true,
		},
		{
			"invalid argument",
			args{zerrors.InvalidArgument("invalid argument")},
			codes.InvalidArgument,
			"invalid argument",
			"",
			true,
		},
		{
			"not found",
			args{zerrors.NotFound("not found")},
			codes.NotFound,
			"not found",
			"",
			true,
		},
		{
			"permission denied",
			args{zerrors.PermissionDenied("permission denied")},
			codes.PermissionDenied,
			"permission denied",
			"",
			true,
		},
		{
			"precondition failed",
			args{zerrors.FailedPrecondition("precondition failed")},
			codes.FailedPrecondition,
			"precondition failed",
			"",
			true,
		},
		{
			"unauthenticated",
			args{zerrors.Unauthenticated("unauthenticated")},
			codes.Unauthenticated,
			"unauthenticated",
			"",
			true,
		},
		{
			"unavailable",
			args{zerrors.Unavailable("unavailable")},
			codes.Unavailable,
			"unavailable",
			"",
			true,
		},
		{
			"unimplemented",
			args{zerrors.Unimplemented("unimplemented")},
			codes.Unimplemented,
			"unimplemented",
			"",
			true,
		},
		{
			"exhausted",
			args{zerrors.ResourceExhausted("exhausted")},
			codes.ResourceExhausted,
			"exhausted",
			"",
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
