package impl

import (
	"context"
	"errors"
	"iter"
	"maps"
	"strings"

	"github.com/hypershadow-io/contract/eb"
	"github.com/hypershadow-io/contract/fmt"
	"github.com/hypershadow-io/contract/json"
	"github.com/hypershadow-io/contract/meta"
)

// Make creates a new generic Builder instance for building structured errors of type E.
// Initializes an empty wrap stack and prepares the builder for fluent-style chaining.
func Make[E error]() Builder[E] { return Builder[E]{wrapped: make([]error, 0, 1)} }

// Wrap wraps the given base error into the provided error builder.
// If the base error is already an eb.Builder, the new builder is added as a wrap to it.
// Otherwise, the base error is added as a wrap to the new builder.
func Wrap(base error, err eb.Builder) eb.Builder {
	if bb, ok := base.(eb.Builder); ok {
		return bb.AddWrap(err)
	}
	return err.AddWrap(base)
}

// Builder is a generic error builder used to construct structured errors of type E.
// Typically used via the [eb.Builder] interface.
type Builder[E error] struct {
	message    string            // human-readable error message
	key        string            // machine-readable error key for client/log correlation
	validation map[string]string // field-level validation errors
	code       int               // numeric error code (e.g. HTTP status)
	wrapped    []error           // additional errors to wrap
	logMessage string            // optional log message (internal use)
	meta       meta.Meta         // attached metadata for extended context
	logger     eb.LogFunc        // optional logger for side-effect logging
}

var noLogger = func(context.Context, string, ...any) {}

func (a Builder[E]) GetMessage() string {
	if a.message != "" {
		return a.message
	}
	for e := range a.unwrap() {
		if result := e.GetMessage(); result != "" {
			return result
		}
	}
	return ""
}
func (a Builder[E]) SetMessagef(format string, args ...any) eb.Builder {
	if len(args) == 0 {
		a.message = format
	} else {
		a.message = fmt.Sprintf(format, args...)
	}
	return a
}

func (a Builder[E]) GetKey() string {
	if a.key != "" {
		return a.key
	}
	for e := range a.unwrap() {
		if result := e.GetKey(); result != "" {
			return result
		}
	}
	return ""
}
func (a Builder[E]) SetKey(v string) eb.Builder {
	a.key = v
	return a
}

func (a Builder[E]) GetValidation() map[string]string {
	if a.validation != nil {
		return a.validation
	}
	for e := range a.unwrap() {
		if result := e.GetValidation(); result != nil {
			return result
		}
	}
	return nil
}
func (a Builder[E]) SetValidation(v map[string]string) eb.Builder {
	a.validation = v
	return a
}

func (a Builder[E]) GetCode() int {
	if a.code != 0 {
		return a.code
	}
	for e := range a.unwrap() {
		if result := e.GetCode(); result != 0 {
			return result
		}
	}
	return 0
}
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

func (a Builder[E]) GetLogMessage() string {
	if a.logMessage != "" {
		return a.logMessage
	}
	for e := range a.unwrap() {
		if result := e.GetLogMessage(); result != "" {
			return result
		}
	}
	return ""
}
func (a Builder[E]) SetLogMessagef(format string, args ...any) eb.Builder {
	if len(args) == 0 {
		a.logMessage = format
	} else {
		a.logMessage = fmt.Sprintf(format, args...)
	}
	return a
}

func (a Builder[E]) GetMeta() meta.Meta {
	if a.meta != nil {
		return a.meta
	}
	for e := range a.unwrap() {
		if result := e.GetMeta(); result != nil {
			return result
		}
	}
	return nil
}
func (a Builder[E]) SetMeta(m meta.Meta) eb.Builder {
	a.meta = m
	return a
}
func (a Builder[E]) MergeMeta(m meta.Meta) eb.Builder {
	a.meta = a.meta.Merge(m)
	return a
}
func (a Builder[E]) DrainMeta() meta.Meta {
	result := meta.Make(len(a.meta))
	for e := range a.unwrap() {
		maps.Copy(result, e.DrainMeta())
	}
	maps.Copy(result, a.meta)
	clear(a.meta)
	return result
}

func (a Builder[E]) GetLogger() eb.LogFunc {
	if a.logger != nil {
		return a.logger
	}
	for e := range a.unwrap() {
		if result := e.GetLogger(); result != nil {
			return result
		}
	}
	return nil
}
func (a Builder[E]) SetLogger(v eb.LogFunc) eb.Builder {
	a.logger = v
	return a
}
func (a Builder[E]) SetNoLogger() eb.Builder {
	return a.SetLogger(noLogger)
}

func (a Builder[E]) Unwrap() []error { return a.wrapped }

func (a Builder[E]) unwrap() iter.Seq[eb.Builder] {
	return func(yield func(eb.Builder) bool) {
		for i := range a.wrapped {
			var e eb.Builder
			if errors.As(a.wrapped[i], &e) {
				if !yield(e) {
					return
				}
			}
		}
	}
}

func (a Builder[E]) Error() string {
	var builder strings.Builder
	var base E
	if any(base) != nil {
		builder.WriteRune('<')
		builder.WriteString(base.Error())
		builder.WriteRune('>')
	}
	if a.message != "" {
		if builder.Len() > 0 {
			builder.WriteRune(' ')
		}
		builder.WriteString(a.message)
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
		if data, _ := json.Marshal(a.meta); len(data) > 0 {
			builder.Write(data)
		}
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
