// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/service/user/v1/user.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateUserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateUserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateUserInfoMultiError,
// or nil if none found.
func (m *CreateUserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NickName

	// no validation rules for Password

	// no validation rules for Mobile

	if len(errors) > 0 {
		return CreateUserInfoMultiError(errors)
	}

	return nil
}

// CreateUserInfoMultiError is an error wrapping multiple validation errors
// returned by CreateUserInfo.ValidateAll() if the designated constraints
// aren't met.
type CreateUserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserInfoMultiError) AllErrors() []error { return m }

// CreateUserInfoValidationError is the validation error returned by
// CreateUserInfo.Validate if the designated constraints aren't met.
type CreateUserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserInfoValidationError) ErrorName() string { return "CreateUserInfoValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserInfoValidationError{}

// Validate checks the field values on UserInfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserInfoResponseMultiError, or nil if none found.
func (m *UserInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Password

	// no validation rules for Mobile

	// no validation rules for NickName

	// no validation rules for Birthday

	// no validation rules for Gender

	// no validation rules for Role

	if len(errors) > 0 {
		return UserInfoResponseMultiError(errors)
	}

	return nil
}

// UserInfoResponseMultiError is an error wrapping multiple validation errors
// returned by UserInfoResponse.ValidateAll() if the designated constraints
// aren't met.
type UserInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserInfoResponseMultiError) AllErrors() []error { return m }

// UserInfoResponseValidationError is the validation error returned by
// UserInfoResponse.Validate if the designated constraints aren't met.
type UserInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoResponseValidationError) ErrorName() string { return "UserInfoResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoResponseValidationError{}

// Validate checks the field values on UserListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserListResponseMultiError, or nil if none found.
func (m *UserListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserListResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserListResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserListResponseMultiError(errors)
	}

	return nil
}

// UserListResponseMultiError is an error wrapping multiple validation errors
// returned by UserListResponse.ValidateAll() if the designated constraints
// aren't met.
type UserListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListResponseMultiError) AllErrors() []error { return m }

// UserListResponseValidationError is the validation error returned by
// UserListResponse.Validate if the designated constraints aren't met.
type UserListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListResponseValidationError) ErrorName() string { return "UserListResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListResponseValidationError{}

// Validate checks the field values on PageInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageInfoMultiError, or nil
// if none found.
func (m *PageInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PageInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pn

	// no validation rules for PSize

	if len(errors) > 0 {
		return PageInfoMultiError(errors)
	}

	return nil
}

// PageInfoMultiError is an error wrapping multiple validation errors returned
// by PageInfo.ValidateAll() if the designated constraints aren't met.
type PageInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageInfoMultiError) AllErrors() []error { return m }

// PageInfoValidationError is the validation error returned by
// PageInfo.Validate if the designated constraints aren't met.
type PageInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageInfoValidationError) ErrorName() string { return "PageInfoValidationError" }

// Error satisfies the builtin error interface
func (e PageInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageInfoValidationError{}

// Validate checks the field values on MobileRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MobileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MobileRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MobileRequestMultiError, or
// nil if none found.
func (m *MobileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MobileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Mobile

	if len(errors) > 0 {
		return MobileRequestMultiError(errors)
	}

	return nil
}

// MobileRequestMultiError is an error wrapping multiple validation errors
// returned by MobileRequest.ValidateAll() if the designated constraints
// aren't met.
type MobileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MobileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MobileRequestMultiError) AllErrors() []error { return m }

// MobileRequestValidationError is the validation error returned by
// MobileRequest.Validate if the designated constraints aren't met.
type MobileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MobileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MobileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MobileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MobileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MobileRequestValidationError) ErrorName() string { return "MobileRequestValidationError" }

// Error satisfies the builtin error interface
func (e MobileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMobileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MobileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MobileRequestValidationError{}

// Validate checks the field values on IdRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdRequestMultiError, or nil
// if none found.
func (m *IdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdRequestMultiError(errors)
	}

	return nil
}

// IdRequestMultiError is an error wrapping multiple validation errors returned
// by IdRequest.ValidateAll() if the designated constraints aren't met.
type IdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdRequestMultiError) AllErrors() []error { return m }

// IdRequestValidationError is the validation error returned by
// IdRequest.Validate if the designated constraints aren't met.
type IdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdRequestValidationError) ErrorName() string { return "IdRequestValidationError" }

// Error satisfies the builtin error interface
func (e IdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdRequestValidationError{}

// Validate checks the field values on UpdateUserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateUserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateUserInfoMultiError,
// or nil if none found.
func (m *UpdateUserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NickName

	// no validation rules for Gender

	// no validation rules for Birthday

	if len(errors) > 0 {
		return UpdateUserInfoMultiError(errors)
	}

	return nil
}

// UpdateUserInfoMultiError is an error wrapping multiple validation errors
// returned by UpdateUserInfo.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserInfoMultiError) AllErrors() []error { return m }

// UpdateUserInfoValidationError is the validation error returned by
// UpdateUserInfo.Validate if the designated constraints aren't met.
type UpdateUserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserInfoValidationError) ErrorName() string { return "UpdateUserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UpdateUserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserInfoValidationError{}

// Validate checks the field values on PasswordCheckInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PasswordCheckInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PasswordCheckInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PasswordCheckInfoMultiError, or nil if none found.
func (m *PasswordCheckInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PasswordCheckInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Password

	// no validation rules for EncryptedPassword

	if len(errors) > 0 {
		return PasswordCheckInfoMultiError(errors)
	}

	return nil
}

// PasswordCheckInfoMultiError is an error wrapping multiple validation errors
// returned by PasswordCheckInfo.ValidateAll() if the designated constraints
// aren't met.
type PasswordCheckInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PasswordCheckInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PasswordCheckInfoMultiError) AllErrors() []error { return m }

// PasswordCheckInfoValidationError is the validation error returned by
// PasswordCheckInfo.Validate if the designated constraints aren't met.
type PasswordCheckInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PasswordCheckInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PasswordCheckInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PasswordCheckInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PasswordCheckInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PasswordCheckInfoValidationError) ErrorName() string {
	return "PasswordCheckInfoValidationError"
}

// Error satisfies the builtin error interface
func (e PasswordCheckInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPasswordCheckInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PasswordCheckInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PasswordCheckInfoValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckResponseMultiError, or
// nil if none found.
func (m *CheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return CheckResponseMultiError(errors)
	}

	return nil
}

// CheckResponseMultiError is an error wrapping multiple validation errors
// returned by CheckResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckResponseMultiError) AllErrors() []error { return m }

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}
