// Package errors provides unified error definitions and error handling utilities.
package errors

import (
	"fmt"

	pkgErrors "github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error codes for different error types
const (
	CodeInvalidArgument = codes.InvalidArgument
	CodeNotFound        = codes.NotFound
	CodeInternal        = codes.Internal
	CodeUnauthenticated = codes.Unauthenticated
	CodeUnavailable     = codes.Unavailable
)

var (
	// ErrInvalidArgument indicates that the request contains invalid arguments
	ErrInvalidArgument = NewError(CodeInvalidArgument, "invalid argument")

	// ErrContractNotFound indicates that the contract address is not found or invalid
	ErrContractNotFound = NewError(CodeNotFound, "contract not found")

	// ErrClientNotInitialized indicates that the Ethereum client is not initialized
	ErrClientNotInitialized = NewError(CodeInternal, "ethereum client not initialized")

	// ErrChainIDNotConfigured indicates that the chain ID is not configured
	ErrChainIDNotConfigured = NewError(CodeInternal, "chain ID not configured")

	// ErrTransactionFailed indicates that a transaction failed
	ErrTransactionFailed = NewError(CodeInternal, "transaction failed")

	// ErrInvalidPrivateKey indicates that the private key is invalid
	ErrInvalidPrivateKey = NewError(CodeInvalidArgument, "invalid private key")

	// ErrInvalidAddress indicates that an Ethereum address is invalid
	ErrInvalidAddress = NewError(CodeInvalidArgument, "invalid address")

	// ErrInvalidAmount indicates that an amount is invalid
	ErrInvalidAmount = NewError(CodeInvalidArgument, "invalid amount")
)

// AppError represents an application error with a gRPC status code
type AppError struct {
	Code    codes.Code
	Message string
	Err     error
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// GRPCStatus returns the gRPC status for this error
func (e *AppError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Error())
}

// NewError creates a new application error
func NewError(code codes.Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// WrapError wraps an existing error with an application error
func WrapError(err error, code codes.Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// ToGRPCError converts an error to a gRPC error
func ToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	// If it's already an AppError, return its gRPC status
	if appErr, ok := err.(*AppError); ok {
		return appErr.GRPCStatus().Err()
	}

	// Check if it's already a gRPC status error
	if _, ok := status.FromError(err); ok {
		return err
	}

	// Wrap with internal error
	return status.New(CodeInternal, err.Error()).Err()
}

// InvalidArgument returns an invalid argument error
func InvalidArgument(format string, args ...interface{}) *AppError {
	return &AppError{
		Code:    CodeInvalidArgument,
		Message: fmt.Sprintf(format, args...),
	}
}

// InternalError returns an internal error
func InternalError(format string, args ...interface{}) *AppError {
	return &AppError{
		Code:    CodeInternal,
		Message: fmt.Sprintf(format, args...),
	}
}

// Wrap wraps an error with additional context
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return pkgErrors.Wrap(err, message)
}

// Wrapf wraps an error with formatted additional context
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return pkgErrors.Wrapf(err, format, args...)
}

