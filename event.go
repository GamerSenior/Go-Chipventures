package main

type eventID int

// Enum contendo todos os possíveis eventos
const (
	MOUSEBUTTONPRESSED eventID = iota
	MOUSEBUTTONRELEASED
	KEYPRESSED
	KEYRELEASED
	WINDOWCLOSE
	WINDOWRESIZE
	WINDOWFOCUS
	WINDOWLOSTFOCUS
)

type eventCategory uint8

// Enum contendo os tipos de eventos
const (
	KEYBOARDEVENT eventCategory = iota
	MOUSEEVENT
	WINDOWEVENT
)

// Event é a estrutura básica de um evento
type Event struct {
	name string
	// id        eventID
	// category  eventCategory
	callbacks []interface{}
}

func (e *Event) fire(data interface{}) {
	for index := 0; index < len(e.callbacks); index++ {
		e.callbacks[index].(func(interface{}))(data)
	}
}

func (e *Event) registerCallback(callback interface{}) {
	e.callbacks = append(e.callbacks, callback)
}

func (e *Event) unregisterCallback(callback interface{}, i int) {
	e.callbacks = append(e.callbacks[:i], e.callbacks[i+1:]...)
}

// Dispatcher é a estrutura que toma conta dos eventos
type Dispatcher struct {
	events map[string]Event
}

func (d *Dispatcher) dispatch(name string, data interface{}) {
	event := d.events[name]
	event.fire(data)
}

func (d *Dispatcher) on(name string, callback interface{}) {
	event := d.events[name]
	if event.name != "" {
		event = Event{name, make([]interface{}, 1)}
	}
	event.registerCallback(callback)
}

func (d *Dispatcher) off(name string, callback interface{}) {
	event := d.events[name]
	if event.name != "" {
		for i := 0; i < len(event.callbacks); i++ {
			if callback == event.callbacks[i] {
				event.unregisterCallback(callback, i)
				if len(event.callbacks) == 0 {
					delete(d.events, name)
				}
			}
		}
	}
}
