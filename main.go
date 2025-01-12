package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("enter your first name:\n")
	var name string
	fmt.Scan(&name)
	fmt.Printf("enter you last name:\n")
	var last_name string
	fmt.Scan(&last_name)
	var isnamevalid bool = len(name) >= 2 && len(last_name) >= 2
	if isnamevalid == true {
		fmt.Printf("Keep goingon %v %v\n", name, last_name)
	}
	if isnamevalid != true {
		fmt.Printf("write a valid first name and last name\n")
		fmt.Printf("%v %v not the valid name\n", name, last_name)
	}
	//cheking the validation for given info is correct or not
	fmt.Printf("enter your E-mail :\n")
	var email string
	fmt.Scan(&email)
	var ismailr bool = strings.Contains(email, "@")
	if ismailr == true {
		fmt.Printf("keep going on %v %v\n", name, last_name)
		fmt.Printf("how is you experience going with us\n")
		fmt.Printf("rate us out of 10\n")
		var rate int
		fmt.Scan(&rate)
		fmt.Printf("* * * * * * * * * *\n")
		for i := 0; i < rate; i++ {
			fmt.Printf("*")
			fmt.Printf(" ")
		}
		fmt.Println()
		fmt.Printf("thank for rate us %v/10", rate)
	}
	if ismailr != true {
		fmt.Printf("gave the correct email adress")
	}

	//switch in go lang
	fmt.Println("enter the coutry name \n")
	var coutryname string
	fmt.Scan(&coutryname)
	city := coutryname

	switch city {
	case "maxico":
		fmt.Printf("welcome to maxico")
	case "london":
		fmt.Printf("welcome to london")
	case "paris":
		fmt.Printf("welcome to paris")
	case "new york":
		fmt.Printf("welcome to new york")

	}

}
