package event

import (
	"core/logger"
	"fmt"
	"runtime/debug"
	"sync"
)

type Event struct {
	Source interface{}
	Topic  string
}

type Channel chan Event

type Listener interface {
	OnListen(Event)
}

type ChannelSlice []Channel

type EventBus struct {
	subscribers map[string]ChannelSlice
	rm          sync.RWMutex
}

func (e *EventBus) Publish(topic string, source interface{}) {
	go func() {
		e.rm.RLock()
		if chans, found := e.subscribers[topic]; found {
			channels := append(ChannelSlice{}, chans...)
			go func(event Event, channelSlices ChannelSlice) {
				for _, ch := range channelSlices {
					ch <- event
				}
			}(Event{Source: source, Topic: topic}, channels)

		}
		e.rm.RUnlock()
	}()
}

func (e *EventBus) Subscribe(topic string, listener Listener) {
	var channel = make(Channel)
	e.rm.Lock()
	if prev, found := e.subscribers[topic]; found {
		e.subscribers[topic] = append(prev, channel)
	} else {
		e.subscribers[topic] = append([]Channel{}, channel)
	}
	e.rm.Unlock()
	go func() {
		for event := range channel {
			func() {
				defer func() {
					if r := recover(); r != nil {
						debug.PrintStack()
						err := fmt.Sprintf("%s", r)
						logger.Error("panic error : %v", err)
					}
				}()
				listener.OnListen(event)
			}()
		}
	}()
}

func (e *EventBus) UnSubscribe(topic string) {
	e.rm.Lock()
	channels, ex := e.subscribers[topic]
	if !ex {
		return
	}
	for _, v := range channels {
		close(v)
	}
	delete(e.subscribers, topic)
	e.rm.Unlock()
}

var (
	Bus = &EventBus{
		subscribers: map[string]ChannelSlice{},
	}
)
