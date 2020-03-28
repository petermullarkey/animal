package main

/* Test the program by running it twice and testing it with a different sequence of integers each time. 
The first test sequence of integers should be all positive numbers 
and the second test should have at least one negative number.
*/

import ("testing"
		"reflect"
)

func Test(t *testing.T) {
     var tests = []struct {
     	 input []string; want string
	}{
		{[]string{"cow", "eat"}, "grass"},
		{[]string{"bird", "speak"}, "peep"},
	}
	var retVal string
	for _, c := range tests {
		retVal = AnimalTrait(c.input[0], c.input[1])
	    if !reflect.DeepEqual(retVal, c.want) {
	       t.Errorf("AnimalTrait(%q) wanted %q", c.input[0], c.want)
	    }
	}
}