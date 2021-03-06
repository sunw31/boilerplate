package event

import (
	fmt "fmt"
	"time"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
)

// [DDD Reference - Domain Language](https://domainlanguage.com/ddd/reference/)
const (
	// Occurred is a tag for the time the event occurred.
	Occurred = "occurred"
	// Entered is a tag for the time the event entered.
	Entered = "entered"
)

var (
	GetTopicKeyValueWithContexter = keyvalue.GetTopicKeyValueWithContexter
)

type (
	KeyValueWithContext   = keyvalue.KeyValueWithContext
	KeyValueWithContexter = keyvalue.KeyValueWithContexter
)

// NewEvent returns a timestamp for the time the event occurred.
func NewEvent(timeStamp *time.Time, origin string) (Eventer, error) {
	const op = op + ".NewEvent"
	if timeStamp.IsZero() {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be given. timeStamp is zero(%v)", timeStamp)}
	}
	if origin == "" {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. origin is ''"}
	}
	v0 := xid.New()
	var v1 Eventer = &Event{
		Id:        v0.Bytes(),
		Origin:    origin,
		TimePoint: make([]*TimePoint, 0),
	}
	v1, err := v1.AddTimePoint(Occurred, *timeStamp)
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, err
	}
	v1, err = v1.AddTimePoint(Entered, v0.Time())
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, err
	}
	return v1, nil
}

// AddTimePoint returns a new Eventer with given the TimePoint.
func (p *Event) AddTimePoint(tag string, timeStamp time.Time) (Eventer, error) {
	const op = op + ".Event.AddTimePoint"
	if tag == "" {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. tag is ''"}
	}
	if timeStamp.IsZero() {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be given. timeStamp is zero(%v)", timeStamp)}
	}
	return &Event{
		Id:     p.GetId(),
		Origin: p.GetOrigin(),
		TimePoint: append(p.GetTimePoint(),
			&TimePoint{
				Timestamp: &timestamp.Timestamp{
					Seconds: timeStamp.Unix(),
					Nanos:   int32(timeStamp.UnixNano()),
				},
				Tag: tag,
			},
		),
	}, nil
}

func (p *Event) GetKey() []byte {
	return []byte(fmt.Sprintf("%s+%s", p.GetOrigin(), p.GetId()))
}

//go:generate interfacer -for github.com/michilu/boilerplate/service/event.Event -as event.Eventer -o entity-Eventer.go

//go:generate genny -in=../topic/with-context.go -out=gen-EventWithContext.go -pkg=$GOPACKAGE gen "T=Event Ier=Eventer"
//go:generate genny -in=../topic/topic.go -out=gen-topic-EventWithContexter.go -pkg=$GOPACKAGE gen "ChanT=EventWithContexter"
//go:generate interfacer -for github.com/michilu/boilerplate/service/event.EventWithContext -as event.EventWithContexter -o if-EventWithContexter.go
