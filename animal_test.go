package main

/* Test the program by running it twice and testing it with a different sequence of integers each time. 
The first test sequence of integers should be all positive numbers 
and the second test should have at least one negative number.
*/

import ("testing"
		"reflect"
		"fmt"
)

func TestNewAnimal(t *testing.T) {
     var tests = []struct {
     	 input []string; want string
	}{
		{[]string{"newanimal", "charlie", "cow"}, "charlie"},
		{[]string{"newanimal", "sammy", "snake"}, "sammy"},
	}
	var retVal userAnimal
	for _, c := range tests {
		retVal = createNewAnimal(c.input[1], c.input[2])
	    if !reflect.DeepEqual(retVal.Name, c.want) {
	       t.Errorf("createNewAnimal(%q) wanted %q", c.input[0], c.want)
	    }
	}
}
func TestAnimalQuery(t *testing.T) {
     var tests = []struct {
     	 input []string; want string
	}{
		{[]string{"query", "charlie", "eat", "cow"}, "grass"},
		{[]string{"query", "sammy", "speak", "snake"}, "hsss"},
	}
	var retVal string
	for _, c := range tests {
		var nA = createNewAnimal(c.input[1], c.input[3])
		fmt.Println("name is " + nA.Name )
		fmt.Println(" type is: " +reflect.TypeOf(nA.Animal).String())
		retVal = AnimalTrait(nA, c.input[2])
	    if !reflect.DeepEqual(retVal, c.want) {
	       t.Errorf("AnimalTrait(%q) wanted %q got %q", c.input[2], c.want, retVal)
	    }
	}
}

func TestSpeak(t *testing.T) {
	var newIAnimal Snake
	newIAnimal.InitMe("mice", "slither", "hsss")		
	fmt.Printf("in test type is %T", newIAnimal )
	// tstAn := userAnimal{"suzz", newIAnimal}
	retVal := newIAnimal.Speak()
	if retVal != "hsss" {
		t.Errorf("should be hsss was %q", retVal)
	}
}