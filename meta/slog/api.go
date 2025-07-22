package slog

import (
	"log/slog"

	"github.com/hypershadow-io/contract/meta"
)

// ToAttrs converts meta.Meta into a slice of slog-compatible attributes.
// Nested meta structures (meta.Meta or map[string]any) are recursively converted into slog groups.
// The result can be passed directly to slog logging methods.
func ToAttrs(m meta.Meta) []any {
	result := make([]any, 0, len(m))
	for k, v := range m {
		switch vTyped := v.(type) {
		case meta.Meta:
			result = append(result, slog.Group(k, ToAttrs(vTyped)...))
		case map[string]any:
			result = append(result, slog.Group(k, ToAttrs(vTyped)...))
		default:
			result = append(result, slog.Any(k, v))
		}
	}
	return result
}
