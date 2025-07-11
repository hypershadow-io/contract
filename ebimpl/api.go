package ebimpl

import (
	"context"
	"errors"
	"strings"

	"github.com/hypershadow-io/contract/eb"
	"github.com/hypershadow-io/contract/fmt"
	"github.com/hypershadow-io/contract/json"
	"github.com/hypershadow-io/contract/meta"
)

// Make creates a new generic Builder instance for building structured errors of type E.
// Initializes an empty wrap stack and prepares the builder for fluent-style chaining.
func Make[E error]() Builder[E] { return Builder[E]{wrapped: make([]error, 0, 1)} }

// Builder is a generic error builder used to construct structured errors of type E.
// Typically used via the [eb.Builder] interface.
type Builder[E error] struct {
	base       error             // base (original) error
	Message    string            `json:"message,omitempty"`    // human-readable error message
	Key        string            `json:"key,omitempty"`        // machine-readable error key for client/log correlation
	Validation map[string]string `json:"validation,omitempty"` // field-level validation errors
	code       int               // numeric error code (e.g. HTTP status)
	wrapped    []error           // additional errors to wrap
	logMessage string            // optional log message (internal use)
	meta       meta.Meta         // attached metadata for extended context
	logger     eb.LogFunc        // optional logger for side-effect logging
}

var noLogger = func(context.Context, string, ...any) {}

func (a Builder[E]) GetBase() error { return a.base }
func (a Builder[E]) SetBase(v error) eb.Builder {
	a.base = v
	return a
}

func (a Builder[E]) GetMessage() string {
	if a.Message != "" {
		return a.Message
	}
	for _, err := range a.wrapped {
		if e, ok := err.(eb.Builder); ok {
			msg := e.GetMessage()
			if msg != "" {
				return msg
			}
		}
	}
	return ""
}

func (a Builder[E]) SetMessagef(format string, args ...any) eb.Builder {
	if len(args) == 0 {
		a.Message = format
	} else {
		a.Message = fmt.Sprintf(format, args...)
	}
	return a
}

func (a Builder[E]) GetKey() string { return a.Key }
func (a Builder[E]) SetKey(v string) eb.Builder {
	a.Key = v
	return a
}

func (a Builder[E]) GetValidation() map[string]string { return a.Validation }
func (a Builder[E]) SetValidation(v map[string]string) eb.Builder {
	a.Validation = v
	return a
}

func (a Builder[E]) GetCode() int { return a.code }
func (a Builder[E]) SetCode(v int) eb.Builder {
	a.code = v
	return a
}

func (a Builder[E]) AddWrap(v error) eb.Builder {
	if v != nil {
		a.wrapped = append(a.wrapped, v)
	}
	return a
}

func (a Builder[E]) GetLogMessage() string { return a.logMessage }
func (a Builder[E]) SetLogMessagef(format string, args ...any) eb.Builder {
	if len(args) == 0 {
		a.logMessage = format
	} else {
		a.logMessage = fmt.Sprintf(format, args...)
	}
	return a
}

func (a Builder[E]) GetMeta() meta.Meta { return a.meta }
func (a Builder[E]) SetMeta(m meta.Meta) eb.Builder {
	a.meta = m
	return a
}
func (a Builder[E]) MergeMeta(m meta.Meta) eb.Builder {
	a.meta = a.meta.Merge(m)
	return a
}

func (a Builder[E]) GetLogger() eb.LogFunc { return a.logger }
func (a Builder[E]) SetLogger(v eb.LogFunc) eb.Builder {
	a.logger = v
	return a
}
func (a Builder[E]) SetNoLogger() eb.Builder {
	return a.SetLogger(noLogger)
}

func (a Builder[E]) Unwrap() []error {
	if any(a.base) == nil {
		return a.wrapped
	}
	return append(
		append(make([]error, 0, 1+len(a.wrapped)), a.base),
		a.wrapped...,
	)
}

func (a Builder[E]) Error() string {
	var builder strings.Builder
	var base E
	errors.As(a.base, &base)
	if any(base) != nil {
		builder.WriteRune('<')
		builder.WriteString(base.Error())
		builder.WriteRune('>')
	}
	if a.Message != "" {
		if builder.Len() > 0 {
			builder.WriteRune(' ')
		}
		builder.WriteString(a.Message)
	}
	if a.logMessage != "" {
		if builder.Len() > 0 {
			builder.WriteRune(' ')
		}
		builder.WriteString("log: ")
		builder.WriteString(a.logMessage)
	}
	if !a.meta.IsZero() {
		if builder.Len() > 0 {
			builder.WriteRune(' ')
		}
		builder.WriteString("meta: ")
		_ = json.NewEncoder(&builder).Encode(a.meta)
	}
	for _, err := range a.wrapped {
		if builder.Len() > 0 {
			builder.WriteString(", err: ")
		}
		builder.WriteString(err.Error())
	}
	if builder.Len() > 0 {
		return builder.String()
	}
	return "unknown"
}
