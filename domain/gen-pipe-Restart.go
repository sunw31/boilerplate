// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package domain

import (
	"context"

	"github.com/michilu/boilerplate/v/errs"
	"google.golang.org/grpc/codes"
)

// GetPipeRestart returns new input(chan<- Durationer)/output(<-chan Struct) channels that embedded the given 'func(Durationer) Struct'.
func GetPipeRestart(
	ctx context.Context,
	fn func(Durationer) (struct{}, error),
	fnErr func(error) bool,
) (
	chan<- Durationer,
	<-chan struct{},
) {
	const op = "pipe.GetPipeRestart"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan Durationer)
	outCh := make(chan struct{})

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				if fnErr(&errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			select {
			case <-ctx.Done():
				err := ctx.Err()
				if err != nil {
					fnErr(&errs.Error{Op: op, Err: err})
				}
				return
			case outCh <- o:
			default:
			}
		}
	}()

	return inCh, outCh
}

// GetFanoutRestart returns new input(chan<- Durationer)/output(<-chan Struct) channels that embedded the given 'func(Durationer) Struct'.
func GetFanoutRestart(
	ctx context.Context,
	fn func(Durationer) ([]struct{}, error),
	fnErr func(error) bool,
) (
	chan<- Durationer,
	<-chan struct{},
) {
	const op = "pipe.GetFanoutRestart"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan Durationer)
	outCh := make(chan struct{})

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				if fnErr(&errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			for _, v := range o {
				select {
				case <-ctx.Done():
					err := ctx.Err()
					if err != nil {
						fnErr(&errs.Error{Op: op, Err: err})
					}
					return
				case outCh <- v:
				default:
				}
			}
		}
	}()

	return inCh, outCh
}