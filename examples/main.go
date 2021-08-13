package main

import (
	"queue/queue"
	"time"
)

func main() {
	q := queue.New()

	_ = q.CreateTopic("A")
	_ = q.CreateTopic("B")

	_ = q.CreateSubscription(NewSub1("sub1topicA"), "A")
	_ = q.CreateSubscription(NewSub1("sub2topicA"), "A")
	_ = q.CreateSubscription(NewSub1("sub2topicB"), "B")

	q.StartSubscriptions()

	q.Publish("A", msg("AAA"))
	q.Publish("A", msg("BBB"))

	q.Publish("B", msg("YOYO"))
	q.Publish("B", msg("XOXO"))

	time.Sleep(10*time.Second)
	q.ResetOffset("sub1topicA")
	time.Sleep(10*time.Second)
	q.ResetOffset("sub2topicB")

	time.Sleep(1111*time.Second)
}


