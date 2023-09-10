package src

import "sync"

var messageBus *MessageBus

type MessageBus struct {
	lock   sync.Mutex
	subs   map[string][]chan<- string
	quit   chan struct{}
	closed bool
}

func NewMessageBus() *MessageBus {
	if messageBus == nil {
		messageBus = &MessageBus{
			subs: make(map[string][]chan<- string),
			quit: make(chan struct{}),
		}
	}

	return messageBus
}

func (b *MessageBus) Publish(topic, message string) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.closed {
		return
	}

	for _, channel := range b.subs[topic] {
		channel <- message
	}
}

func (b *MessageBus) Subscribe(topic string, channel chan<- string) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.closed {
		return
	}

	b.subs[topic] = append(b.subs[topic], channel)
}

func (b *MessageBus) Close() {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.closed {
		return
	}

	b.closed = true
	close(b.quit)

	for _, channel := range b.subs {
		for _, sub := range channel {
			close(sub)
		}
	}
}
