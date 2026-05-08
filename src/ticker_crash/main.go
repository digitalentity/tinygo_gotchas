package main

import (
	"machine"
	"runtime"
	"time"
)

func led1() {
	led := machine.LED_GREEN

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)
		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}

func led2() {
	led := machine.LED_RED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	t := time.NewTicker(time.Millisecond * 500)
	s := bool(false)
	for range t.C {
		if s {
			led.Low()
		} else {
			led.High()
		}
		s = !s
	}
}

func main() {
	go led1()
	go led2()

	for {
		runtime.Gosched()
	}
}
