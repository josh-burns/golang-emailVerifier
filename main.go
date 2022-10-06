package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	email := os.Args[1]

	allowedDomains := []string{"net", "com"}
	illegalCharacters := []string{"¬", "*", "^", "£", "$", "!"}
	split := strings.Split(email, "@")

	switch length := len(split); {
	case length > 2:
		fmt.Println("****** FAIL: there is no @ sign")
		return
	case length > 3:
		fmt.Println("****** FAIL: there are too many @ signs")
		return
	default:
		fmt.Println("processing...")
	}

	beforeAmpersand, afterAmpersand := split[0], split[1]

	afterAmpersandSplit := strings.Split(afterAmpersand, ".")

afterAmp:
	for _, a := range allowedDomains {
		if afterAmpersandSplit[1] == a {
			break afterAmp
		} else {
			fmt.Println("****** FAIL: Email does not comply with domain rule, must be: ", allowedDomains)
			return
		}
	}

	for _, a := range illegalCharacters {
		splitChars := strings.Split(beforeAmpersand, "")
		for _, b := range splitChars {
			if a == b {
				fmt.Println("****** FAIL: Forbidden Character: ", b)
				return
			}
		}
	}

	fmt.Println("_________PASS_______", email, "is VALID")
}
