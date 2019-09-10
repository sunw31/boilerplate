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
	topicByteWithContexter *mapByteWithContexter
)

func init() {
	topicByteWithContexter = newMapByteWithContexter()
}

// GetTopicByteWithContexter returns a TopicByteWithContexter of the given topic.
func GetTopicByteWithContexter(topic interface{}) TopicByteWithContexter {
	return topicByteWithContexter.get(topic)
}

// TopicByteWithContexter is a topic.
type TopicByteWithContexter interface {
	// Publish returns a '<-chan ByteWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan ByteWithContexter)
	// Publisher returns a 'chan<- ByteWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- ByteWithContexter
	// Subscribe returns a 'chan<- ByteWithContexter' that joins to the given topic.
	Subscribe(c chan<- ByteWithContexter)
}

type tByteWithContexter struct {
	mu sync.RWMutex
	c  []chan<- ByteWithContexter
}

func newTByteWithContexter() *tByteWithContexter {
	return &tByteWithContexter{
		c: make([]chan<- ByteWithContexter, 0),
	}
}

func (t *tByteWithContexter) Publish(ctx context.Context, c <-chan ByteWithContexter) {
	const op = op + ".tByteWithContexter.Publish"

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
				go func(c chan<- ByteWithContexter, v ByteWithContexter) {
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

func (t *tByteWithContexter) Publisher(ctx context.Context) chan<- ByteWithContexter {
	const op = op + ".tByteWithContexter.Publisher"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}

	c := make(chan ByteWithContexter)
	t.Publish(ctx, c)
	return c
}

func (t *tByteWithContexter) Subscribe(c chan<- ByteWithContexter) {
	const op = op + ".tByteWithContexter.Subscribe"

	if c == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'c' is nil"})
	}

	t.mu.Lock()
	t.c = append(t.c, c)
	t.mu.Unlock()
}

type mapByteWithContexter struct {
	mu sync.RWMutex
	m  map[interface{}]*tByteWithContexter
}

func newMapByteWithContexter() *mapByteWithContexter {
	return &mapByteWithContexter{
		m: make(map[interface{}]*tByteWithContexter),
	}
}

func (m *mapByteWithContexter) get(topic interface{}) TopicByteWithContexter {
	const op = op + ".mapByteWithContexter.get"

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
	v = newTByteWithContexter()
	m.m[topic] = v
	return v
}