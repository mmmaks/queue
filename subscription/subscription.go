package subscription

import (
	"queue/message"
	"time"
)

type Subscription interface {
	Consume(message message.Message)
	Delay() time.Duration
	Name() string
}

