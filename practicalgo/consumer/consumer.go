package main

import "fmt"

type Consumer struct {
	name string
	age  int
}

type Consumers []Consumer

func (c Consumers) YoungConsumer() Consumers {
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
			name: "mike",
			age:  31,
		},
		Consumer{
			name: "tom",
			age:  23,
		},
		Consumer{
			name: "ai",
			age:  17,
		},
	}

	for _, c := range consumers.YoungConsumer() {
		fmt.Printf("Name: %s, Age: %d\n", c.name, c.age)
	}
}
