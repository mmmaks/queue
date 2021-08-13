package subscription

import (
	"queue/topic"
	"sync"
)

type Manager interface {
	CreateSubscription(Subscription, string) error
	StartSubscriptions()
	ResetOffset(subName string)
}

type manager struct {
	subscriptionConsumers map[string]Consumer
	topicManager topic.Manager
	mu sync.Mutex
}

func NewManager(topicManager topic.Manager) Manager {
	return &manager{topicManager:topicManager, mu: sync.Mutex{}, subscriptionConsumers: map[string]Consumer{}}
}

func (m *manager) CreateSubscription(sub Subscription, topicName string) error {
	defer m.synchronize()()
	m.subscriptionConsumers[sub.Name()] = NewConsumer(sub, topicName, m.topicManager)
	return nil
}

func (m *manager) StartSubscriptions() {
	for _, v := range m.subscriptionConsumers {
		v.Consume()
	}
}

func (m *manager) ResetOffset(subName string) {
	m.subscriptionConsumers[subName].Reset()
}


func (m *manager) synchronize() func() {
	m.mu.Lock()
	return m.mu.Unlock
}