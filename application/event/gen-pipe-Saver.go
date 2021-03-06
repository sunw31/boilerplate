// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package event

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/event"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type SaverGetContexter interface {
	GetContext() context.Context
}

// GetPipeSaver returns new input(chan<- EventKeyValueWithContexter)/output(<-chan ContextContext) channels that embedded the given 'func(EventKeyValueWithContexter) ContextContext'.
func GetPipeSaver(
	ctx context.Context,
	fn func(event.KeyValueWithContexter) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.KeyValueWithContexter,
	<-chan context.Context,
) {
	const op = op + ".GetPipeSaver"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan event.KeyValueWithContexter)
	outCh := make(chan context.Context)

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v, ok := i.(SaverGetContexter)
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

// GetFanoutSaver returns new input(chan<- EventKeyValueWithContexter)/output(<-chan ContextContext) channels that embedded the given 'func(EventKeyValueWithContexter) ContextContext'.
func GetFanoutSaver(
	ctx context.Context,
	fn func(event.KeyValueWithContexter) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.KeyValueWithContexter,
	<-chan context.Context,
) {
	const op = op + ".GetFanoutSaver"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan event.KeyValueWithContexter)
	outCh := make(chan context.Context)

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var ctx0 context.Context
				ctx0, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(SaverGetContexter)
					if ok {
						ctx0 = v1.GetContext()
					} else {
						ctx0 = context.Background()
					}
				}
				v2 := fnErr(ctx0, &errs.Error{Op: op, Err: err})
				v0 := trace.FromContext(ctx0)
				if v0 != nil {
					v0.End()
				}
				if v2 {
					return
				}
				continue
			}
			for _, v3 := range v0 {
				select {
				case <-ctx.Done():
					err := ctx.Err()
					if err != nil {
						fnErr(ctx, &errs.Error{Op: op, Err: err})
					}
					return
				case outCh <- v3:
				default:
				}
			}
		}
	}()

	return inCh, outCh
}
