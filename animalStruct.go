package main

import (
	"fmt"
)

type Animal struct { 
	eat string
	move string
	speak string
}
func (an *Animal) InitMe(eat, move, speak string) {
	an.eat = eat
	an.move = move
	an.speak = speak
}
func (an *Animal) Eat() string {
	fmt.Println(an.eat)
	return an.eat
}
func (an *Animal) Move() string {
	fmt.Println(an.move)
	return an.move
}	
func (an *Animal) Speak() string {
	fmt.Println(an.speak)
	return an.speak
}