// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: message.proto

package event

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetId()) < 1 {
		return EventValidationError{
			field:  "Id",
			reason: "value length must be at least 1 bytes",
		}
	}

	if utf8.RuneCountInString(m.GetOrigin()) < 1 {
		return EventValidationError{
			field:  "Origin",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetTimePoint()) < 1 {
		return EventValidationError{
			field:  "TimePoint",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTimePoint() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  fmt.Sprintf("TimePoint[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

// Validate checks the field values on TimePoint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TimePoint) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimePointValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetTag()) < 1 {
		return TimePointValidationError{
			field:  "Tag",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// TimePointValidationError is the validation error returned by
// TimePoint.Validate if the designated constraints aren't met.
type TimePointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimePointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimePointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimePointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimePointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimePointValidationError) ErrorName() string { return "TimePointValidationError" }

// Error satisfies the builtin error interface
func (e TimePointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimePoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimePointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimePointValidationError{}