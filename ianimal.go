package main

/*  Define an interface type called Animal which describes the methods of an animal. 
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
 which take no arguments and return no values. 
*/

type iAnimal interface {
	Eat() string
	Move() string
	Speak() string
}
