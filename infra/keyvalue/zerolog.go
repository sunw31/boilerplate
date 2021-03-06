package keyvalue

import (
	"encoding/json"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

func (p *Key) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".Key"
	v, err := json.Marshal(&p)
	if err != nil {
		const op = op + ".MarshalZerologObject.json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return
	}
	e.RawJSON(op, v)
}

func (p *KeyValue) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".KeyValue"
	v, err := json.Marshal(&p)
	if err != nil {
		const op = op + ".MarshalZerologObject.json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return
	}
	e.RawJSON(op, v)
}

func (p *Prefix) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".Prefix"
	v, err := json.Marshal(&p)
	if err != nil {
		const op = op + ".MarshalZerologObject.json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return
	}
	e.RawJSON(op, v)
}
