// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package keyvalue

import (
	"context"
	"fmt"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

// KeyWithContext is Key with context.Context.
type KeyWithContext struct {
	Context context.Context
	Key     Keyer
}

// GetContext returns context.Context.
func (p *KeyWithContext) GetContext() context.Context {
	return p.Context
}

// GetKey returns Keyer.
func (p *KeyWithContext) GetKey() Keyer {
	return p.Key
}

// MarshalZerologObject writes KeyWithContext to given zerolog.Event.
func (p *KeyWithContext) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".KeyWithContext.MarshalZerologObject"
	if p.Key == nil {
		return
	}
	v, ok := p.Key.(zerolog.LogObjectMarshaler)
	if !ok {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument,
			Message: "'*KeyWithContext.Key' must be zerolog.LogObjectMarshaler'"}
		slog.Logger().Error().Str("op", op).Err(err).Msg(err.Error())
		return
	}
	e.Object("KeyWithContext", v)
}

// String returns KeyWithContext as string.
func (p *KeyWithContext) String() string {
	const v0 = "KeyWithContext<Context: %v, Key: %v>"
	if p.Key == nil {
		return fmt.Sprintf(v0, p.Context, p.Key)
	}
	v1, ok := p.Key.(fmt.Stringer)
	if !ok {
		return fmt.Sprintf(v0, p.Context, p.Key)
	}
	return fmt.Sprintf(v0, p.Context, v1.String())
}

// Validate returns error if failed validate.
func (p *KeyWithContext) Validate() error {
	const op = op + ".KeyWithContext.Validate"
	if p.Context == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*KeyWithContext.Context' is nil"}
		return err
	}
	if p.Key == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*KeyWithContext.Key' is nil"}
		return err
	}
	v0, ok := p.Key.(interface{ Validate() error })
	if !ok {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "'*KeyWithContext.Key' must be have 'Validate() error'"}
		return err
	}
	{
		err := v0.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			return err
		}
	}
	return nil
}
