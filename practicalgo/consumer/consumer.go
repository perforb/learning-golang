package main

import (
	"fmt"
	"time"
)

type Consumer struct {
	name      string
	age       int
	expiredAt int64
}

type Consumers []Consumer

func (c Consumers) expires(now time.Time) Consumers {
	consumers := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.expiredAt < now.UnixMilli() {
			consumers = append(consumers, v)
		}
	}
	return consumers
}

func (c Consumers) ExpiredConsumers() Consumers {
	return c.YoungConsumers().expires(time.Now())
}

func (c Consumers) YoungConsumers() Consumers {
	consumers := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.age < 25 {
			consumers = append(consumers, v)
		}
	}
	return consumers
}

func main() {
	consumers := Consumers{
		Consumer{
			name:      "mike",
			age:       31,
			expiredAt: time.Now().Add(1 * time.Minute).UnixMilli(),
		},
		Consumer{
			name:      "tom",
			age:       23,
			expiredAt: time.Now().Add(-1 * time.Minute).UnixMilli(),
		},
		Consumer{
			name:      "ai",
			age:       17,
			expiredAt: time.Now().Add(1 * time.Minute).UnixMilli(),
		},
	}

	for _, c := range consumers.ExpiredConsumers() {
		fmt.Printf("Name: %s, Age: %d, ExpiredAt: %d\n",
			c.name, c.age, c.expiredAt)
	}
}
