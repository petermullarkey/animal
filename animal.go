package main

/*Write a program which allows the user to create a set of animals and to get information about those animals. 
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create 
a new animal of one of the three types, or the user can request information about an animal that he/she has 
already created. Each animal has a unique name, defined by the user. Note that the user can define animals 
of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table 
contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. 
Your program should accept one command at a time from the user, print out a response, and print out a new prompt 
on a new line. Your program should continue in this loop forever. 
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. 
The second string is an arbitrary string which will be the name of the new animal. The third string is the type 
of the new animal, either “cow”, “bird”, or “snake”. Your program should process each newanimal command by 
creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
 The second string is the name of the animal. The third string is the name of the information requested 
 about the animal, either “eat”, “move”, or “speak”. Your program should process each query command 
 by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface 
should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. 
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, 
and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. 
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake 
all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. 
Your program should call the appropriate method when the user issues a query command.
*/
   
import (
	"fmt"
)

func main() {
	var cmdType, name, typeOrBehavior string
	
	currentAnimals := make([]userAnimal, 3)

	for {
		fmt.Printf("Enter newanimal, followed by the new name, and then the animal type: either cow, bird, or snake, OR query, followed by the animal name, and then one of eat, move, or speak (seperated by spaces)> ")
	n, err := fmt.Scanf("%s%s%s", &cmdType, &name, &typeOrBehavior)
	fmt.Printf("Got %d characters\n", n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("here's the input [cmd: " + cmdType + " typeOrbehavior: " + typeOrBehavior + "]")
	if (cmdType == "end" || cmdType == "quit") {
		break
	}
	if (cmdType == "newanimal") {
		currentAnimals = append(currentAnimals, createNewAnimal(name, typeOrBehavior))
	} else if (cmdType == "query") {
		for _, v := range currentAnimals {
			if v.Name == name {
				AnimalTrait(v, typeOrBehavior)
			}
		}
	}
}
}

func createNewAnimal (newName, animalType string) userAnimal {

	if animalType == "cow" {
		var newIAnimal Cow
		newIAnimal.InitMe("grass", "walk", "moo")
		fmt.Printf("in createNewAnimal type is %T", newIAnimal )
		return userAnimal{newName, newIAnimal}
	} else if animalType == "bird" {
		var newIAnimal Bird
		newIAnimal.InitMe("worms", "fly",	"peep")
		return userAnimal{newName, newIAnimal}
	} else if animalType == "snake" {
		var newIAnimal Snake
		newIAnimal.InitMe("mice", "slither", "hsss")		
		fmt.Printf("in createNewAnimal type is %T", newIAnimal )
		tstAn := userAnimal{newName, newIAnimal}
		fmt.Println("in createNewAnimal, snake speak is:  " + AnimalTrait(tstAn, "speak"))
		return tstAn
	} else {
		var newIAnimal Animal
		newIAnimal.InitMe("food", "locomote", "noise")
		return userAnimal{newName, newIAnimal}
	}
}

func AnimalTrait(animalInstance userAnimal, trait string) (returnForTesting string) {
	var thisIAnimal iAnimal
	thisIAnimal = animalInstance.Animal
	fmt.Printf("in AnimalTrait type is %T", thisIAnimal )
	if (trait == "eat") {
		fmt.Println(" in AnimalTrait trait is: " + trait)
		returnForTesting = thisIAnimal.Eat()
	} else if (trait == "move") {
		returnForTesting = thisIAnimal.Move()
	} else if (trait == "speak") {
		fmt.Println(" in AnimalTrait trait is: " + trait)
		returnForTesting = thisIAnimal.Speak()
	} else {
		fmt.Println(" No traits matched by [" + trait + "]")
	}
	return returnForTesting
}
