package Middleware

import (
	"context"
	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
)

// @url https://github.com/go-playground/mold/blob/master/_examples/full/main.go
var conform = modifiers.New()

func SanitizeRegisterTag(tag string, fn mold.Func) {
	conform.Register(tag, fn)
}

func Sanitize(ctx context.Context, v any) error {
	return conform.Struct(ctx, v)
}
