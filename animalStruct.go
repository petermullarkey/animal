package main

/* Define three types Cow, Bird, and Snake. 
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake 
all satisfy the Animal interface.   */

import (
	"fmt"
)

type Cow struct { 
	eat string
	move string
	speak string
}
func (an *Cow) InitMe(eat, move, speak string) {
	an.eat = eat
	an.move = move
	an.speak = speak
}
func (an Cow) Eat() string {
	if &an == nil {
		fmt.Println("generic food")
		return "generic food"
	} else {
	fmt.Println(an.eat)
	return an.eat
	}
}
func (an Cow) Move() string {
	fmt.Println(an.move)
	return an.move
}	
func (an Cow) Speak() string {
	fmt.Println(an.speak)
	return an.speak
}

type Snake struct { 
	eat string
	move string
	speak string
}
func (an *Snake) InitMe(eat, move, speak string) {
	an.eat = eat
	an.move = move
	an.speak = speak
}
func (an Snake) Eat() string {
	fmt.Println(an.eat)
	return an.eat
}
func (an Snake) Move() string {
	fmt.Println(an.move)
	return an.move
}	
func (an Snake) Speak() string {
	if &an == nil {
		fmt.Println("generic noise")
		return "generic noise"
	} else {
		fmt.Println(an.speak)
		return an.speak
	}
}

type Bird struct { 
	eat string
	move string
	speak string
}
func (an *Bird) InitMe(eat, move, speak string) {
	an.eat = eat
	an.move = move
	an.speak = speak
}
func (an Bird) Eat() string {
	fmt.Println(an.eat)
	return an.eat
}
func (an Bird) Move() string {
	fmt.Println(an.move)
	return an.move
}	
func (an Bird) Speak() string {
	fmt.Println(an.speak)
	return an.speak
}

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
func (an Animal) Eat() string {
	fmt.Println(an.eat)
	return an.eat
}
func (an Animal) Move() string {
	fmt.Println(an.move)
	return an.move
}	
func (an Animal) Speak() string {
	fmt.Println(an.speak)
	return an.speak
}

type userAnimal struct { 
	Name string
	Animal iAnimal
}
