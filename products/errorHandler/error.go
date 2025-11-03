package errorhandler

import "github.com/vektah/gqlparser/v2/gqlerror"

// Base constructor
func New(code, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
		Extensions: map[string]any{
			"code": code,
		},
	}
}
