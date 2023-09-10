package main

import (
	"fmt"
	"time"

	"github.com/joaomisturini/contabilista/src"
)

var topic string = "sample-topic"

func main() {
	bus := src.NewMessageBus()

	myChannel := make(chan string)

	bus.Subscribe(topic, myChannel)
	fmt.Println("subscribed")

	go publishSomething(bus)
	fmt.Println("called publishSomething")

	for message := range myChannel {
		fmt.Printf("Your message: %s\n", message)
	}
}

func publishSomething(bus *src.MessageBus) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)

		fmt.Println("published message")
		bus.Publish(topic, fmt.Sprintf("Test message %d", i))
	}

	bus.Close()
}
