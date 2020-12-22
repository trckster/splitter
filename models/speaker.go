package models

import "fmt"

type Speaker struct {
	Word string
}

func NewSpeaker(word string) *Speaker {
	return &Speaker{word}
}

func (speaker *Speaker) Say() {
	fmt.Printf(speaker.Word)
}