package main

import (
	"fmt"
	"queue/message"
	"queue/subscription"
	"time"
)

type sub1 struct {
	name string
}

func (s sub1) Consume(message message.Message) {
	fmt.Printf(" %v message received: %v\n", s.name, message.Body())
	time.Sleep(1*time.Second)
	fmt.Printf(" %v message consumed: %v\n", s.name, message.Body())
}

func (s sub1) Delay() time.Duration {
	return 1*time.Second
}

func (s sub1) Name() string {
	return s.name
}

func NewSub1(name string) subscription.Subscription {
	return &sub1{name: name}
}


type msg string

func (m msg) Body() interface{} {
	return m
}