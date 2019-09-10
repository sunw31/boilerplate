// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package event

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

type StartGetContexter interface {
	GetContext() context.Context
}

// GetPipeStart returns new input(chan<- ContextContext)/output(<-chan EventWithContexter) channels that embedded the given 'func(ContextContext) EventWithContexter'.
func GetPipeStart(
	ctx context.Context,
	fn func(context.Context) (EventWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan EventWithContexter,
) {
	const op = op + ".GetPipeStart"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan context.Context)
	outCh := make(chan EventWithContexter)

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v, ok := i.(StartGetContexter)
					if ok {
						vctx = v.GetContext()
					} else {
						vctx = context.Background()
					}
				}
				if fnErr(vctx, &errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			select {
			case <-ctx.Done():
				err := ctx.Err()
				if err != nil {
					fnErr(ctx, &errs.Error{Op: op, Err: err})
				}
				return
			case outCh <- o:
			default:
			}
		}
	}()

	return inCh, outCh
}

// GetFanoutStart returns new input(chan<- ContextContext)/output(<-chan EventWithContexter) channels that embedded the given 'func(ContextContext) EventWithContexter'.
func GetFanoutStart(
	ctx context.Context,
	fn func(context.Context) ([]EventWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan EventWithContexter,
) {
	const op = op + ".GetFanoutStart"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan context.Context)
	outCh := make(chan EventWithContexter)

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v, ok := i.(StartGetContexter)
					if ok {
						vctx = v.GetContext()
					} else {
						vctx = context.Background()
					}
				}
				if fnErr(vctx, &errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			for _, v := range o {
				select {
				case <-ctx.Done():
					err := ctx.Err()
					if err != nil {
						fnErr(ctx, &errs.Error{Op: op, Err: err})
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
