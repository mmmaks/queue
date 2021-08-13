package topic

import (
	"errors"
	"queue/message"
	"sync"
)

type Manager interface {
	CreateTopic(string) error
	Publish(topic string, message message.Message)
	NextIterator(topic string, it int) (int, error)
	NewIterator() int
	Message(topic string, iterator int) message.Message
}

type manager struct {
	topics map[string]Topic
	mu sync.Mutex
}

func NewManager() Manager {
	return &manager{mu: sync.Mutex{}, topics: map[string]Topic{}}
}

func (m *manager) CreateTopic(s string) error {

	defer m.synchronize()()
	_, ok := m.topics[s]
	if ok {
		return errors.New("topic name already exists")
	}
	m.topics[s] = New(s)
	return nil
}

func (m *manager) Publish(topic string, message message.Message) {
	m.topics[topic].Publish(message)
}

func (m *manager) NextIterator(topic string, it int) (int, error) {
	return m.topics[topic].NextIterator(it)
}

func (m *manager) NewIterator() int {
	return -1
}

func (m *manager) Message(topic string, iterator int) message.Message {
	return m.topics[topic].Message(iterator)
}

func (m *manager) synchronize() func() {
	m.mu.Lock()
	return m.mu.Unlock
}