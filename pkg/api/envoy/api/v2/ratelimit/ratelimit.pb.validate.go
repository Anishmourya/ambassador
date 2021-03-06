// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/ratelimit/ratelimit.proto

package envoy_api_v2_ratelimit

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on RateLimitDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitDescriptor) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetEntries()) < 1 {
		return RateLimitDescriptorValidationError{
			field:  "Entries",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RateLimitDescriptorValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// RateLimitDescriptorValidationError is the validation error returned by
// RateLimitDescriptor.Validate if the designated constraints aren't met.
type RateLimitDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptorValidationError) ErrorName() string {
	return "RateLimitDescriptorValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptorValidationError{}

// Validate checks the field values on RateLimitDescriptor_Entry with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitDescriptor_Entry) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetKey()) < 1 {
		return RateLimitDescriptor_EntryValidationError{
			field:  "Key",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetValue()) < 1 {
		return RateLimitDescriptor_EntryValidationError{
			field:  "Value",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RateLimitDescriptor_EntryValidationError is the validation error returned by
// RateLimitDescriptor_Entry.Validate if the designated constraints aren't met.
type RateLimitDescriptor_EntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptor_EntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptor_EntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptor_EntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptor_EntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptor_EntryValidationError) ErrorName() string {
	return "RateLimitDescriptor_EntryValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptor_EntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor_Entry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptor_EntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptor_EntryValidationError{}
