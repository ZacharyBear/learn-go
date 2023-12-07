package main

import (
	"github.com/google/wire"
)

func InitializeEvent(phrase string) (Event, error) {
	// woops! NewEventNumber is unused
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
