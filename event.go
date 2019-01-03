package main

import "fmt"

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

// TODO implementar os enums definidos anteriormente para
// filtrar melhor os tipos de eventos

// Event é a estrutura básica de um evento
type Event struct {
	Name string
	// id        eventID
	// category  eventCategory
	Callbacks []interface{}
}

// NewEvent inicializa um novo evento
func NewEvent(name string) (e Event) {
	e.Name = name
	e.Callbacks = make([]interface{}, 0)
	return
}

func (e *Event) fire(data ...interface{}) {
	for index := 0; index < len(e.Callbacks); index++ {
		e.Callbacks[index].(func(...interface{}))(data...)
	}
}

func (e *Event) registerCallback(callback interface{}) {
	e.Callbacks = append(e.Callbacks, callback)
	fmt.Println("Callback registrada")
}

func (e *Event) unregisterCallback(callback interface{}, i int) {
	e.Callbacks = append(e.Callbacks[:i], e.Callbacks[i+1:]...)
}

// Dispatcher é a estrutura que toma conta dos eventos
type Dispatcher struct {
	events map[string]Event
}

// NewDispatcher instancia um novo dispatcher
func NewDispatcher() (d Dispatcher) {
	d.events = make(map[string]Event)
	return
}

func (d *Dispatcher) dispatch(name string, data ...interface{}) {
	fmt.Println("Emitindo evento [", name, "] -> ", data)
	event, ok := d.events[name]
	if ok {
		event.fire(data...)
	}
}

func (d *Dispatcher) on(name string, callback interface{}) {
	event, ok := d.events[name]
	if !ok {
		fmt.Println("Evento não existe, criando novo evento...")
		event = NewEvent(name)
	}
	fmt.Println("Registrando callback do evento: ", name)
	event.registerCallback(callback)
	d.events[name] = event
}

func (d *Dispatcher) off(name string, callback interface{}) {
	event, ok := d.events[name]
	if ok {
		for i := 0; i < len(event.Callbacks); i++ {
			if callback == event.Callbacks[i] {
				event.unregisterCallback(callback, i)
				if len(event.Callbacks) == 0 {
					delete(d.events, name)
				}
			}
		}
	}
}
