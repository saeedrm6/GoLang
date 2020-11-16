package structures

import "fmt"

func Types() {
	/*
		its better to define them out of functions ( after package name )

		type YourTypeName struct{
			field1 data_type
			field2 data_type
			.
			.
			.
		}
	*/

	type animal struct {
		name            string
		age             int
		characteristics map[string]string
	}

	var a animal
	a.name = "ship"
	options := make(map[string]string)
	options["sound"] = "baaaa baaa"
	options["color"] = "brown/white"
	a.characteristics = options
	a.age = 4

	fmt.Println("Animal : ", a.name, "Model : ", a)

}
