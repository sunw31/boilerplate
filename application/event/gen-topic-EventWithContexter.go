// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package event

import (
	"context"
	"sync"

	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

var (
	topicEventWithContexter *mapEventWithContexter
)

func init() {
	topicEventWithContexter = newMapEventWithContexter()
}

// GetTopicEventWithContexter returns a TopicEventWithContexter of the given topic.
func GetTopicEventWithContexter(topic interface{}) TopicEventWithContexter {
	return topicEventWithContexter.get(topic)
}

// TopicEventWithContexter is a topic.
type TopicEventWithContexter interface {
	// Publish returns a '<-chan EventWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan EventWithContexter)
	// Publisher returns a 'chan<- EventWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- EventWithContexter
	// Subscribe returns a 'chan<- EventWithContexter' that joins to the given topic.
	Subscribe(c chan<- EventWithContexter)
}

type tEventWithContexter struct {
	mu sync.RWMutex
	c  []chan<- EventWithContexter
}

func newTEventWithContexter() *tEventWithContexter {
	return &tEventWithContexter{
		c: make([]chan<- EventWithContexter, 0),
	}
}

func (t *tEventWithContexter) Publish(ctx context.Context, c <-chan EventWithContexter) {
	const op = op + ".tEventWithContexter.Publish"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	go func() {
	loop:
		select {
		case <-ctx.Done():
			return
		default:
		}
		for v := range c {
			for _, c := range t.c {
				go func(c chan<- EventWithContexter, v EventWithContexter) {
					select {
					case <-ctx.Done():
						return
					case c <- v:
					}
				}(c, v)
			}
			goto loop
		}
	}()

}

func (t *tEventWithContexter) Publisher(ctx context.Context) chan<- EventWithContexter {
	const op = op + ".tEventWithContexter.Publisher"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}

	c := make(chan EventWithContexter)
	t.Publish(ctx, c)
	return c
}

func (t *tEventWithContexter) Subscribe(c chan<- EventWithContexter) {
	const op = op + ".tEventWithContexter.Subscribe"

	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	t.mu.Lock()
	t.c = append(t.c, c)
	t.mu.Unlock()
}

type mapEventWithContexter struct {
	mu sync.RWMutex
	m  map[interface{}]*tEventWithContexter
}

func newMapEventWithContexter() *mapEventWithContexter {
	return &mapEventWithContexter{
		m: make(map[interface{}]*tEventWithContexter),
	}
}

func (m *mapEventWithContexter) get(topic interface{}) TopicEventWithContexter {
	const op = op + ".mapEventWithContexter.get"

	if topic == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'topic' is nil"})
	}

	m.mu.RLock()
	v, ok := m.m[topic]
	m.mu.RUnlock()
	if ok {
		return v
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.m[topic]
	if ok {
		return v
	}
	v = newTEventWithContexter()
	m.m[topic] = v
	return v
}
