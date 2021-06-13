// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ocp-note-api.proto

package ocp_note_api

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
)

// Validate checks the field values on CreateNoteV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateNoteV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateNoteV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetClassroomId() <= 0 {
		return CreateNoteV1RequestValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetDocumentId() <= 0 {
		return CreateNoteV1RequestValidationError{
			field:  "DocumentId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateNoteV1RequestValidationError is the validation error returned by
// CreateNoteV1Request.Validate if the designated constraints aren't met.
type CreateNoteV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteV1RequestValidationError) ErrorName() string {
	return "CreateNoteV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteV1RequestValidationError{}

// Validate checks the field values on CreateNoteV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateNoteV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NoteId

	return nil
}

// CreateNoteV1ResponseValidationError is the validation error returned by
// CreateNoteV1Response.Validate if the designated constraints aren't met.
type CreateNoteV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteV1ResponseValidationError) ErrorName() string {
	return "CreateNoteV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteV1ResponseValidationError{}

// Validate checks the field values on MultiCreateNotesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateNotesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNotes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateNotesV1RequestValidationError{
					field:  fmt.Sprintf("Notes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateNotesV1RequestValidationError is the validation error returned by
// MultiCreateNotesV1Request.Validate if the designated constraints aren't met.
type MultiCreateNotesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateNotesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateNotesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateNotesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateNotesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateNotesV1RequestValidationError) ErrorName() string {
	return "MultiCreateNotesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateNotesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateNotesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateNotesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateNotesV1RequestValidationError{}

// Validate checks the field values on MultiCreateNotesV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateNotesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NumberOfNotesCreated

	return nil
}

// MultiCreateNotesV1ResponseValidationError is the validation error returned
// by MultiCreateNotesV1Response.Validate if the designated constraints aren't met.
type MultiCreateNotesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateNotesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateNotesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateNotesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateNotesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateNotesV1ResponseValidationError) ErrorName() string {
	return "MultiCreateNotesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateNotesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateNotesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateNotesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateNotesV1ResponseValidationError{}

// Validate checks the field values on DescribeNoteV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeNoteV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNoteId() <= 0 {
		return DescribeNoteV1RequestValidationError{
			field:  "NoteId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeNoteV1RequestValidationError is the validation error returned by
// DescribeNoteV1Request.Validate if the designated constraints aren't met.
type DescribeNoteV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeNoteV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeNoteV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeNoteV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeNoteV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeNoteV1RequestValidationError) ErrorName() string {
	return "DescribeNoteV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeNoteV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeNoteV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeNoteV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeNoteV1RequestValidationError{}

// Validate checks the field values on DescribeNoteV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeNoteV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeNoteV1ResponseValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeNoteV1ResponseValidationError is the validation error returned by
// DescribeNoteV1Response.Validate if the designated constraints aren't met.
type DescribeNoteV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeNoteV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeNoteV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeNoteV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeNoteV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeNoteV1ResponseValidationError) ErrorName() string {
	return "DescribeNoteV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeNoteV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeNoteV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeNoteV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeNoteV1ResponseValidationError{}

// Validate checks the field values on ListNotesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListNotesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() <= 0 {
		return ListNotesV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	if m.GetOffset() < 0 {
		return ListNotesV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// ListNotesV1RequestValidationError is the validation error returned by
// ListNotesV1Request.Validate if the designated constraints aren't met.
type ListNotesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotesV1RequestValidationError) ErrorName() string {
	return "ListNotesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotesV1RequestValidationError{}

// Validate checks the field values on ListNotesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListNotesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNotes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListNotesV1ResponseValidationError{
					field:  fmt.Sprintf("Notes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListNotesV1ResponseValidationError is the validation error returned by
// ListNotesV1Response.Validate if the designated constraints aren't met.
type ListNotesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotesV1ResponseValidationError) ErrorName() string {
	return "ListNotesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotesV1ResponseValidationError{}

// Validate checks the field values on RemoveNoteV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveNoteV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNoteId() <= 0 {
		return RemoveNoteV1RequestValidationError{
			field:  "NoteId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveNoteV1RequestValidationError is the validation error returned by
// RemoveNoteV1Request.Validate if the designated constraints aren't met.
type RemoveNoteV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveNoteV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveNoteV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveNoteV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveNoteV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveNoteV1RequestValidationError) ErrorName() string {
	return "RemoveNoteV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveNoteV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveNoteV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveNoteV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveNoteV1RequestValidationError{}

// Validate checks the field values on RemoveNoteV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveNoteV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveNoteV1ResponseValidationError is the validation error returned by
// RemoveNoteV1Response.Validate if the designated constraints aren't met.
type RemoveNoteV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveNoteV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveNoteV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveNoteV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveNoteV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveNoteV1ResponseValidationError) ErrorName() string {
	return "RemoveNoteV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveNoteV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveNoteV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveNoteV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveNoteV1ResponseValidationError{}

// Validate checks the field values on Note with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Note) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for ClassroomId

	// no validation rules for DocumentId

	return nil
}

// NoteValidationError is the validation error returned by Note.Validate if the
// designated constraints aren't met.
type NoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteValidationError) ErrorName() string { return "NoteValidationError" }

// Error satisfies the builtin error interface
func (e NoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteValidationError{}

// Validate checks the field values on NewNote with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *NewNote) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return NewNoteValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetClassroomId() <= 0 {
		return NewNoteValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetDocumentId() <= 0 {
		return NewNoteValidationError{
			field:  "DocumentId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// NewNoteValidationError is the validation error returned by NewNote.Validate
// if the designated constraints aren't met.
type NewNoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewNoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewNoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewNoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewNoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewNoteValidationError) ErrorName() string { return "NewNoteValidationError" }

// Error satisfies the builtin error interface
func (e NewNoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewNoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewNoteValidationError{}
