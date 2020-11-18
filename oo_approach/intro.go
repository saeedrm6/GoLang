package oo_approach

func Help() {
	/*
		Is Golang OO?  Yes and No
		What is Object Oriented Programming?
			○ is a programming paradigm based on the concept of “objects”,
			  which may contain data, in the form of fields, often known as attributes;
			  and code, in the form of procedures, often known as methods ○ ○ an Object’s procedures
			  can access and often modify the attributes of the object with which they are associated
			○ an Object’s internal state is protected from outside world (encapsulated) leveraging
			   private/protected/public visibility of attributes and methods
			○ an Object is frequently defined in OO languages as an instance of a Class
	*/

	/*
		How Languages Implement it?
		○ Encapsulation (possible on package level in Go)
		○ Composition (possible through embedding in Go)
		○ Polymorphism (possible through Interface satisfaction in Go)
		○ Inheritance (Go does not provide)
	*/

	/*
		Struct in Structures Sections is a best model to start
		struct is like class
		Variables name in struct if start with Capital it means that field is public
	*/

	type Creature struct {
		Name string
		Real bool
	}

	/*
		Name & Real are public field
	*/
}

/*
types must be define outside of functions to be accessible in same package go files
 */
type Creature struct {
	Name string
	Real bool
}