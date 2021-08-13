package topic

import (
	"errors"
	"queue/message"
	"sync"
)

type Topic interface {
	Publish(message message.Message)
	Message(iterator int) message.Message
	NextIterator(it int) (int, error)
}

type topic struct {
	name string
	message []message.Message
	mu sync.Mutex
}

func New(name string) *topic {
	return &topic{name: name, mu: sync.Mutex{}}
}

func (t *topic) Publish(message message.Message) {
	defer t.synchronize()()
	t.message = append(t.message, message)
}

func (t *topic) Message(iterator int) message.Message {
	return t.message[iterator]
}

func (t *topic) NextIterator(it int) (int, error) {
	if len(t.message) <= it+1 {
		return 0, errors.New("end of list")
	}
	return it+1, nil
}

func (t *topic) Name() string {
	return t.name
}

func (t *topic) synchronize() func() {
	t.mu.Lock()
	return t.mu.Unlock
}