package main

import (
	"log"
	"os"

	"github.com/gregdel/pushover"
	"github.com/warthog618/gpio"
)

func sendToPushover(contents string) {
	app := pushover.New(os.Getenv("PUSHOVER_TOKEN"))
	recipient := pushover.NewRecipient(os.Getenv("PUSHOVER_RECIPIENT"))

	message := pushover.NewMessage(contents)
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
}

func handleDoorOpen(pin *gpio.Pin) {
	sendToPushover("Just letting you know that your door has been opened")
}

func handleDoorClosed(pin *gpio.Pin) {
	sendToPushover("You can relax, the door has now been closed")
}

func main() {
	pin := gpio.NewPin(gpio.J8p7)
	pin.Watch(gpio.EdgeFalling, handleDoorOpen)  // high to low
	pin.Watch(gpio.EdgeRising, handleDoorClosed) // low to high
}
