package main

import "./models"

func main() {
	var s = models.NewSpeaker("Some information")

	s.Say()
}