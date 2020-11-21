package oo_approach

import (
	"fmt"
	"strconv"
)

type NationalID_neww struct {
	StringFormat string
	DigitFormat  [10]int
}

/*
	important note : values of array return in decimal value. you have
	to read http://www.asciitable.com and check decimal of target values
	OR convert your cells of array to string
*/

/*
	your object from NationalID_neww is id. you can fetch data from id object
*/

func (id NationalID_neww) isValidNationalCard3() bool {
	if len(id.StringFormat) == 10 {
		for _, value := range id.StringFormat {
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

var placeValuee = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

// residual constant based on : algorithm
const residualConss = 11
const controlDigiIndexx = 9

func (id NationalID_neww) isCheckSumValid() bool {
	var (
		sum   int
		resid int
	)
	for index := 0; index < 9; index++ {
		digit, err := strconv.Atoi(string(id.StringFormat[index]))
		if err != nil {
			return false
		}
		sum = sum + digit*placeValuee[index]
	}

	resid = sum % residualConss
	controlDigit, err := strconv.Atoi(string(id.StringFormat[controlDigiIndexx]))
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

func Example3() {
	var nID1 NationalID_neww
	nID1.StringFormat = "0123456789"
	fmt.Printf("my national id num %s is valid : %v\n",nID1.StringFormat,nID1.isValidNationalCard3())
}
