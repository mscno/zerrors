package http

import (
	"errors"
	"fmt"
	"github.com/mscno/zerrors"
	"net/http"
	"testing"
)

func TestZitadelErrorToHTTPStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantOk         bool
	}{
		{
			name: "no error",
			args: args{
				err: nil,
			},
			wantStatusCode: http.StatusOK,
			wantOk:         true,
		},
		{
			name: "wrapped already exists",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToAlreadyExists(nil, "id", "message")),
			},
			wantStatusCode: http.StatusConflict,
			wantOk:         true,
		},
		{
			name: "wrapped deadline exceeded",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToDeadlineExceeded(nil, "id", "message")),
			},
			wantStatusCode: http.StatusGatewayTimeout,
			wantOk:         true,
		},
		{
			name: "wrapped internal",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToInternal(nil, "id", "message")),
			},
			wantStatusCode: http.StatusInternalServerError,
			wantOk:         true,
		},
		{
			name: "wrapped invalid argument",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToInvalidArgument(nil, "id", "message")),
			},
			wantStatusCode: http.StatusBadRequest,
			wantOk:         true,
		},
		{
			name: "wrapped not found",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToNotFound(nil, "id", "message")),
			},
			wantStatusCode: http.StatusNotFound,
			wantOk:         true,
		},
		{
			name: "wrapped permission denied",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToPermissionDenied(nil, "id", "message")),
			},
			wantStatusCode: http.StatusForbidden,
			wantOk:         true,
		},
		{
			name: "wrapped precondition failed",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToFailedPrecondition(nil, "id", "message")),
			},
			wantStatusCode: http.StatusBadRequest,
			wantOk:         true,
		},
		{
			name: "wrapped unauthenticated",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToUnauthenticated(nil, "id", "message")),
			},
			wantStatusCode: http.StatusUnauthorized,
			wantOk:         true,
		},
		{
			name: "wrapped unavailable",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToUnavailable(nil, "id", "message")),
			},
			wantStatusCode: http.StatusServiceUnavailable,
			wantOk:         true,
		},
		{
			name: "wrapped unimplemented",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToUnimplemented(nil, "id", "message")),
			},
			wantStatusCode: http.StatusNotImplemented,
			wantOk:         true,
		},
		{
			name: "wrapped resource exhausted",
			args: args{
				err: fmt.Errorf("wrapped %w", zerrors.ToResourceExhausted(nil, "id", "message")),
			},
			wantStatusCode: http.StatusTooManyRequests,
			wantOk:         true,
		},
		{
			name: "no caos/zitadel error",
			args: args{
				err: errors.New("error"),
			},
			wantStatusCode: http.StatusInternalServerError,
			wantOk:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatusCode, gotOk := ZitadelErrorToHTTPStatusCode(tt.args.err)
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("ZitadelErrorToHTTPStatusCode() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
			if gotOk != tt.wantOk {
				t.Errorf("ZitadelErrorToHTTPStatusCode() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
