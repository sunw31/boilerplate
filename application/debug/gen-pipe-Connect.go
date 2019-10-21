// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package debug

import (
	"context"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type ConnectGetContexter interface {
	GetContext() context.Context
}

// GetPipeConnect returns new input(chan<- DebugClientWithContexter)/output(<-chan ContextContext) channels that embedded the given 'func(DebugClientWithContexter) ContextContext'.
func GetPipeConnect(
	ctx context.Context,
	fn func(debug.ClientWithContexter) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- debug.ClientWithContexter,
	<-chan context.Context,
) {
	const op = op + ".GetPipeConnect"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan debug.ClientWithContexter)
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
					v, ok := i.(ConnectGetContexter)
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

// GetFanoutConnect returns new input(chan<- DebugClientWithContexter)/output(<-chan ContextContext) channels that embedded the given 'func(DebugClientWithContexter) ContextContext'.
func GetFanoutConnect(
	ctx context.Context,
	fn func(debug.ClientWithContexter) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- debug.ClientWithContexter,
	<-chan context.Context,
) {
	const op = op + ".GetFanoutConnect"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan debug.ClientWithContexter)
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
					v1, ok := i.(ConnectGetContexter)
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
