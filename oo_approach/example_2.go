package oo_approach

import (
	"fmt"
	"strconv"
)

type NationalID_new string

/*
	important note : values of array return in decimal value. you have
	to read http://www.asciitable.com and check decimal of target values
	OR convert your cells of array to string
*/

/*
	your object from NationalID_new is id. you can fetch data from id object
*/

func (id NationalID_new) isValidNationalCard2() bool {
	if len(id) == 10 {
		for _, value := range id {
			_, err := strconv.Atoi(string(value))
			if err != nil {
				return false
			}
		}
		if id.isCheckSumValid() {
			return true
		}
	}
	return false
}

/*
	Algorithm : http://www.aliarash.com/article/codemeli/codemeli.htm
*/

var placeValue = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

// residual constant based on : algorithm
const residualCons = 11
const controlDigiIndex = 9

func (id NationalID_new) isCheckSumValid() bool {
	var (
		sum   int
		resid int
	)
	for index := 0; index < 9; index++ {
		digit, err := strconv.Atoi(string(id[index]))
		if err != nil {
			return false
		}
		sum = sum + digit*placeValue[index]
	}

	resid = sum % residualCons
	controlDigit, err := strconv.Atoi(string(id[controlDigiIndex]))
	if err != nil {
		return false
	}

	if resid < 2 {
		if resid == controlDigit {
			return true
		}
	} else {
		if controlDigit == (11 - resid) {
			return true
		}
	}
	return false
}

func Example2() {
	var test NationalID_new = "0120456788"
	fmt.Printf("My National ID %s is Valid : %t", test, test.isValidNationalCard2())
}
