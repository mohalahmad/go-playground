package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
	// Or you can define it like this, but the first way is more common
	// contact   contactInfo
}

func main() {
	// how to define a struct
	//
	//-- option 1
	// alex := person{"alex", "anderson"}
	// fmt.Println("-- option 1")
	// fmt.Println(alex)
	//
	//-- option 2
	// alex := person{firstName: "alex", lastName: "anderson"}
	// fmt.Println("-- option 2")
	// fmt.Println(alex)
	//
	//-- option 3
	fmt.Println("-- option 3")
	var alex person
	alex.firstName = "alex"
	alex.lastName = "anderson"
	fmt.Printf("%+v", alex)
	fmt.Println()

	// ---------------------------------

	jim := person{
		firstName: "jim",
		lastName:  "party",
		// contact: contactInfo{  // use this if you defined contactInfo like this: contact   contactInfo
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// the short way to call the updateName function is like this, but it will not modify the original jim struct, it will create a copy of jim and modify the copy, so we need to use a pointer to modify the original jim struct, this is because the updateName function has a pointer receiver "*", it allows us to modify the original struct
	// jim.updateName("jimmy")
	// jim.print()

	// the long way to write the updateName function is like this:
	jimPointer := &jim             // this creates a pointer to the jim struct
	jimPointer.updateName("jimmy") // this is how you call the updateName function using the pointer, it will modify the original jim struct
	jim.print()
	fmt.Println()
	fmt.Println("jim: ", jim)
	fmt.Println("jimPointer: ", jimPointer)
	fmt.Printf("address:  %p\n", jimPointer)
	println("address: ", jimPointer) // this will print the address of jimPointer, which is the same as the address of jim because jimPointer is a pointer to jim
	//println("address: ", jim)  // this will fail to compile because jim is not a pointer, it is a struct, so you cannot print its address using the "%p" format specifier, you need to use the "&" operator to get the address of jim, like this: println("address: ", &jim) or fmt.Printf("address: %p\n", &jim)

	//----------------------------------
	// note that slices dont require pointers to modify the original slice, because slices are reference types, they already contain a pointer to the underlying array, so when you pass a slice to a function, you are passing a copy of the slice header, but the slice header contains a pointer to the underlying array, so when you modify the slice in the function, you are modifying the underlying array, which is shared between all copies of the slice header, so you dont need to use a pointer to modify the original slice, but if you want to modify the slice header itself, like changing the length or capacity of the slice, then you need to use a pointer to modify the original slice header.
	// here an example of modifying a slice without using a pointer:
	mySlice := []int{1, 2, 3}
	fmt.Println("before updateSlice: ", mySlice)
	updateSlice(mySlice) // this will modify the original mySlice because slices are reference types, they already contain a pointer to the underlying array, so when you pass a slice to a function, you are passing a copy of the slice header, but the slice header contains a pointer to the underlying array, so when you modify the slice in the function, you are modifying the underlying array, which is shared between all copies of the slice header, so you dont need to use a pointer to modify the original slice
	fmt.Println("after updateSlice: ", mySlice)
	println("mySlice address: ", mySlice) // this will succeed because mySlice is a slice, it is a reference type, it already contains a pointer to the underlying array, so when you print mySlice, you are printing the slice header, which contains the pointer to the underlying array, so you can print the address of the underlying array using the "%p" format specifier, like this: fmt.Printf("mySlice address: %p\n", mySlice) or println("mySlice address: ", mySlice)
}

// the short way to write this function is like this:
// func (p *person) updateName(newFirstName string) { // this is a pointer receiver "*", it allows us to modify the original struct , this is the short way to write it, the long way described below
// 	p.firstName = newFirstName
// }

// the long way to write this function is like this:
// this function can also used for the short way, but the long way is more explicit and easier to understand for beginners, it shows that we are using a pointer receiver and we are modifying the original struct, while the short way is more concise and easier to read for experienced developers, it is a matter of preference and readability
func (pointerToPerson *person) updateName(newFirstName string) { // this is a pointer receiver "*", it allows us to modify the original struct
	(*pointerToPerson).firstName = newFirstName // this is how you access the struct fields using the pointer, you need to dereference the pointer using "*"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateSlice(s []int) {
	s[0] = 100 // this will modify the original slice because slices are reference types, they already contain a pointer to the underlying array, so when you pass a slice to a function, you are passing a copy of the slice header, but the slice header contains a pointer to the underlying array, so when you modify the slice in the function, you are modifying the underlying array, which is shared between all copies of the slice header, so you dont need to use a pointer to modify the original slice
}
