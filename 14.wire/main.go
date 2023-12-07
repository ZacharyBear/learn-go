package main

import "learn-wire/entity"

func main() {
	message := entity.NewMessage()
	greeter := entity.NewGreeter(message)
	event := entity.NewEvent(greeter)

	event.Start()
}
