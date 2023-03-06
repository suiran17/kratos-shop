// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/goods/v1/goods.proto

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

// Validate checks the field values on CategoryInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryInfoRequestMultiError, or nil if none found.
func (m *CategoryInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for ParentCategory

	// no validation rules for Level

	// no validation rules for IsTab

	// no validation rules for Sort

	if len(errors) > 0 {
		return CategoryInfoRequestMultiError(errors)
	}

	return nil
}

// CategoryInfoRequestMultiError is an error wrapping multiple validation
// errors returned by CategoryInfoRequest.ValidateAll() if the designated
// constraints aren't met.
type CategoryInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryInfoRequestMultiError) AllErrors() []error { return m }

// CategoryInfoRequestValidationError is the validation error returned by
// CategoryInfoRequest.Validate if the designated constraints aren't met.
type CategoryInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryInfoRequestValidationError) ErrorName() string {
	return "CategoryInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryInfoRequestValidationError{}

// Validate checks the field values on CategoryInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryInfoResponseMultiError, or nil if none found.
func (m *CategoryInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for ParentCategory

	// no validation rules for Level

	// no validation rules for IsTab

	// no validation rules for Sort

	if len(errors) > 0 {
		return CategoryInfoResponseMultiError(errors)
	}

	return nil
}

// CategoryInfoResponseMultiError is an error wrapping multiple validation
// errors returned by CategoryInfoResponse.ValidateAll() if the designated
// constraints aren't met.
type CategoryInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryInfoResponseMultiError) AllErrors() []error { return m }

// CategoryInfoResponseValidationError is the validation error returned by
// CategoryInfoResponse.Validate if the designated constraints aren't met.
type CategoryInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryInfoResponseValidationError) ErrorName() string {
	return "CategoryInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryInfoResponseValidationError{}

// Validate checks the field values on GoodsTypeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GoodsTypeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GoodsTypeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GoodsTypeRequestMultiError, or nil if none found.
func (m *GoodsTypeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GoodsTypeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetName()) < 3 {
		err := GoodsTypeRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTypeCode()) < 3 {
		err := GoodsTypeRequestValidationError{
			field:  "TypeCode",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NameAlias

	// no validation rules for IsVirtual

	// no validation rules for Desc

	// no validation rules for Sort

	if len(errors) > 0 {
		return GoodsTypeRequestMultiError(errors)
	}

	return nil
}

// GoodsTypeRequestMultiError is an error wrapping multiple validation errors
// returned by GoodsTypeRequest.ValidateAll() if the designated constraints
// aren't met.
type GoodsTypeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GoodsTypeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GoodsTypeRequestMultiError) AllErrors() []error { return m }

// GoodsTypeRequestValidationError is the validation error returned by
// GoodsTypeRequest.Validate if the designated constraints aren't met.
type GoodsTypeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoodsTypeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoodsTypeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoodsTypeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoodsTypeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoodsTypeRequestValidationError) ErrorName() string { return "GoodsTypeRequestValidationError" }

// Error satisfies the builtin error interface
func (e GoodsTypeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoodsTypeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoodsTypeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoodsTypeRequestValidationError{}

// Validate checks the field values on GoodsTypeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GoodsTypeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GoodsTypeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GoodsTypeResponseMultiError, or nil if none found.
func (m *GoodsTypeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GoodsTypeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GoodsTypeResponseMultiError(errors)
	}

	return nil
}

// GoodsTypeResponseMultiError is an error wrapping multiple validation errors
// returned by GoodsTypeResponse.ValidateAll() if the designated constraints
// aren't met.
type GoodsTypeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GoodsTypeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GoodsTypeResponseMultiError) AllErrors() []error { return m }

// GoodsTypeResponseValidationError is the validation error returned by
// GoodsTypeResponse.Validate if the designated constraints aren't met.
type GoodsTypeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoodsTypeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoodsTypeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoodsTypeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoodsTypeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoodsTypeResponseValidationError) ErrorName() string {
	return "GoodsTypeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GoodsTypeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoodsTypeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoodsTypeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoodsTypeResponseValidationError{}