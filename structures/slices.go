package structures

import "fmt"

func Slices() {
	// is like array but you dont need to specify the length off slice on defining !
	// you need to call make function in first step
	var a = []int{1, 2}
	s := []int{1, 2, 3, 4}
	fmt.Println(s, a, &s, &a)

	// read from start index to end index
	fmt.Println(s[0:2])

	// Error : out of range
	// fmt.Println(s[0:6])

	// Read from custom to end
	fmt.Println(s[2:])

	/*
		SLICE CRUD
	*/

	// Add
	s = append(s, 5)
	s = append(s, 6, 7)
	fmt.Println(s)

	// GET
	fmt.Println(s[2])

	// Update
	s[2] = 124
	fmt.Println(s)

	// Delete
	// target index : 3
	s = append(s[:3], s[3+1:]...)
	fmt.Println(s)

	/*
		Loop Through Slice
	*/

	// For
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// Range
	for key, value := range s {
		fmt.Println(key, value)
	}

}
