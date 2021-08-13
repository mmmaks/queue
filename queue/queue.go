package queue

import (
	"queue/message"
	"queue/subscription"
	"queue/topic"
)

type Queue interface {
	CreateTopic(string) error
	CreateSubscription(subscription.Subscription, string) error
	StartSubscriptions()
	Publish(string, message.Message)
	ResetOffset(subName string)
}

type queue struct {
	topicsManager       topic.Manager
	subscriptionManager subscription.Manager
}

func New() Queue {
	q := &queue{topicsManager: topic.NewManager()}
	q.subscriptionManager = subscription.NewManager(q.topicsManager)
	return q
}

func (q queue) CreateTopic(s string) error {
	return q.topicsManager.CreateTopic(s)
}

func (q queue) CreateSubscription(sub subscription.Subscription, topicName string) error {
	return q.subscriptionManager.CreateSubscription(sub, topicName)
}

func (q queue) StartSubscriptions() {
	q.subscriptionManager.StartSubscriptions()
}

func (q queue) Publish(topic string, m message.Message) {
	q.topicsManager.Publish(topic, m)
}

func (q queue) ResetOffset(subName string) {
	q.subscriptionManager.ResetOffset(subName)
}
