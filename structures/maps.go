package structures

import "fmt"

func Maps() {
	// map[KeyType]ValueType
	// KeyType : Comparable Type(No Slices,Map,Functions)   || example : string,int,float
	// ValueType: Any Type  || example : string
	// map is kind of pointers

	abstraction := make(map[string]int)
	abstraction["weight"] = 95
	abstraction["height"] = 175

	// len() : counting length of map
	fmt.Println(abstraction, abstraction["weight"], len(abstraction))

	var addressbook map[string]string     // this is first step of defining map
	addressbook = make(map[string]string) // start to make and get space in memory , without this you cant assigning data to it
	addressbook["firstName"] = "Saeed"
	addressbook["lastName"] = "Rahimi Manesh"
	fmt.Println(addressbook)

	// delete(map_variable,keyname)  :  if you want to delete custom key
	delete(addressbook, "firstName")
	fmt.Println(addressbook)

	// if you want to make sure from making space for map
	if addressbook == nil {
		fmt.Println("You didnt call make for addressbook map")
	}

	// change & add value of key
	addressbook["firstName"] = "Amir"
	fmt.Println(addressbook)

	// Access item
	z := addressbook["firstName"]
	fmt.Println(z)

	// Access Not-Exist Item
	j := addressbook["mobile"]
	fmt.Println(j) // 0

	// Check & Get
	i, ok := addressbook["mobile"]
	fmt.Println(i, ok)

	// Only check
	_, cok := addressbook["mobile"]
	fmt.Println(cok)

	// Iterate Over Content
	for key, value := range addressbook {
		fmt.Println("key :", key, " & value :", value)
	}

	// Multiple Initialize
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gi":  1908,
		"adg": 912,
	}
	fmt.Println(commits)

	//  Empty Initialize of map
	var empty = map[string]int{}
	fmt.Println(empty)
}
