package oo_approach

import (
	"fmt"
	"strconv"
)

type NationalID string

/*
	important note : values of array return in decimal value. you have
	to read http://www.asciitable.com and check decimal of target values
	OR convert your cells of array to string
*/

func (id NationalID) isValidNationalCard(nID NationalID) bool {
	if len(nID) == 10 {
		for _, value := range nID {
			//fmt.Println(key, value)  => value in decimal\

			/*
				First Method :
				if !(value >= 48 && value <= 57) {
					return false
				}
			*/

			_, err := strconv.Atoi(string(value))
			if err != nil {
				return false
			}

		}
		return true
	}
	return false
}

func Example1() {
	var test NationalID = "012-456789"
	fmt.Printf("My National ID %s is Valid : %t", test, test.isValidNationalCard(test))
}
