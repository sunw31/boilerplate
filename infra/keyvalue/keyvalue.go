package keyvalue

import (
	"context"
)

const (
	op = "infra/keyvalue"
)

type KeyValueCloser interface {
	Close() error
	//Delete(context.Context, keyvalue.Keyer) error
	Get(context.Context, Keyer) (KeyValuer, error)
	Save(context.Context, KeyValuer) error
}
