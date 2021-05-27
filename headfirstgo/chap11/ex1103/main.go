package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
	// n.Walk()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func main() {
	var toy NoiseMaker = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	fmt.Println("-----------------")

	play(Whistle("Toyco Carary"))
	play(Horn("Toyco Blaster"))
	play(Robot("Botco Ambler"))
	fmt.Println("----------------------")
	toy = Robot("Botco Ambler")
	toy.MakeSound()
	var robot Robot = toy.(Robot)
	robot.Walk()

}
