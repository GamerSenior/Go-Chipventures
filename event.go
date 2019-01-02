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

type keyPressedEvent struct {
	keyEvent
}

func NewKeyPressedEvent(code int) keyPressedEvent {
	return keyPressedEvent{
		keyEvent{code},
	}
}
