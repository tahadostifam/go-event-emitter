package event_emitter

import (
	"errors"
)

var (
	ErrEventNotFound = errors.New("event not found")
)

type EventEmitter interface {
	Emit(eventName string, args interface{}) error
	On(eventName string, listener func(args interface{})) error
}

type eventEmitter struct {
	events map[string]func(args interface{})
}

func NewEventEmitter() EventEmitter {
	return &eventEmitter{
		events: make(map[string]func(args interface{})),
	}
}

func (e *eventEmitter) Emit(eventName string, args interface{}) error {
	emit := e.events[eventName]
	if emit != nil {
		emit(args)
		return nil
	}

	return ErrEventNotFound
}

func (e *eventEmitter) On(eventName string, listener func(args interface{})) error {
	event := e.events[eventName]
	if event != nil {
		return ErrEventNotFound
	}

	e.events[eventName] = listener
	return nil
}
