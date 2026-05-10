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
		time.Sleep(time.Millisecond * 200)
		led.High()
		time.Sleep(time.Millisecond * 800)
	}
}

func led2() {
	led := machine.LED_BLUE
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	t := time.NewTicker(time.Millisecond * 20)
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

func led3() {
	led := machine.LED_RED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	t := time.NewTicker(time.Millisecond * 499)
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

func printstat() {
	ms := runtime.MemStats{}
	for {
		runtime.ReadMemStats(&ms)
		println("Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	println("Starting!")

	// Disable NVIC UART interrupt
	machine.DefaultUART.Interrupt.Disable()

	go led1()
	go led2()
	go led3()
	// go printstat()

	for {
		runtime.Gosched()
	}
}
