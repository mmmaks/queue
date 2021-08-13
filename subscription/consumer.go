package subscription

import (
	"fmt"
	"queue/topic"
	"sync"
	"time"
)

type Consumer interface {
	Consume()
	Reset()
}

type consumer struct {
	sub          Subscription
	topic        string
	topicManager topic.Manager
	iterator     int
	mu sync.Mutex
}

func NewConsumer(sub Subscription, topic string, topicManager topic.Manager) Consumer {
	return &consumer{
		sub: sub, topicManager: topicManager,
		topic: topic, iterator: topicManager.NewIterator(),
		mu: sync.Mutex{}}
}

func (c *consumer) Consume() {
	go func() {
		for {
			unlock := c.synchronize()
			it, err := c.topicManager.NextIterator(c.topic, c.iterator)
			if err != nil {
				//fmt.Println("err for it", c.iterator)
				unlock()
				<- time.After(c.sub.Delay())
				continue
			}
			c.iterator = it
			unlock()
			message := c.topicManager.Message(c.topic, c.iterator)
			c.sub.Consume(message)
		}
	}()
}

func (c *consumer) Reset() {
	defer c.synchronize()()
	c.iterator = c.topicManager.NewIterator()
	fmt.Println("resetting iterator for subscription ", c.sub.Name())
}

func (c *consumer) synchronize() func() {
	c.mu.Lock()
	return c.mu.Unlock
}