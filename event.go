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

// Ievent é a interface que que todos os eventos devem implementar
type Ievent interface {
	Log() string
}

type event struct {
	id       eventID
	category eventCategory
	name     string
	handled  bool
}

type keyEvent struct {
	keyCode int
	event
}

// KeyPressedEvent é uma struct contendo um keyEvent
type KeyPressedEvent struct {
	keyEvent
}

// NewKeyPressedEvent é um contrutor que instancia um KeyPressedEvent
func NewKeyPressedEvent(code int) KeyPressedEvent {
	return KeyPressedEvent{
		keyEvent{code, event{KEYPRESSED, KEYBOARDEVENT, "KeyPressedEvent", false}},
	}
}
