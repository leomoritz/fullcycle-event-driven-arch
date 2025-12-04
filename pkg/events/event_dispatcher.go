package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered for this event")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface // Map of event names to their handlers {name: [handlers]}
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
