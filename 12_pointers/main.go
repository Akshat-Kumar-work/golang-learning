package main

import "fmt"

/*
===========================================
UNDERSTANDING * and & IN GO
===========================================

& (address-of operator):
  - Gets the memory address of a variable
  - Example: &num gives you the address where num is stored
  - Use when: You want to pass a variable by reference

* (dereference operator):
  - Gets the VALUE stored at a memory address
  - Example: *ptr gives you the value that ptr points to
  - Use when: You have a pointer and want to access/modify the value

*Type (pointer type):
  - Declares a type that holds a memory address
  - Example: *int means "pointer to int", *Person means "pointer to Person"
  - Use in: Function parameters, return types, variable declarations

===========================================
PRODUCTION PATTERNS WITH STRUCTS
===========================================

1. Function parameter: func updateUser(u *User)
   - *User means "accept a pointer to User"
   - Inside function: u.Name = "new" (Go auto-dereferences)

2. Return type: func NewUser() *User
   - *User means "return a pointer to User"
   - Return: return &User{...} (return address)

3. Variable: var user *User
   - *User means "variable that holds address of User"
   - Assign: user = &User{...} or user = existingUserPtr

4. Method receiver: func (u *User) Update()
   - *User means "method works on pointer to User"
   - Call: user.Update() (Go automatically passes &user)

===========================================
WHY USE POINTERS WITH STRUCTS?
===========================================
- Avoid copying large structs (performance)
- Modify the original struct (not a copy)
- Share the same struct instance across functions
- Return structs from functions efficiently
*/

// ============ BASIC EXAMPLES ============

func changeNumbyvalue(num int) {
	// num is a COPY of the original
	// Changes here don't affect the original
	num = 5
	fmt.Println("in changenum by value", num)
}

func changeNumByReference(num *int) {
	// num is a POINTER (memory address)
	// *num dereferences to get the actual value
	// Changing *num changes the ORIGINAL variable
	*num = 5
	fmt.Println("in changenum by reference (address)", num)
	fmt.Println("in changenum by reference (value)", *num)
}

// ============ STRUCT EXAMPLES (PRODUCTION PATTERNS) ============

type User struct {
	ID    int
	Name  string
	Email string
}

// Pattern 1: Function that accepts pointer to struct
// *User in parameter = "accept a pointer to User"
// This allows modifying the original struct
func updateUserEmail(user *User, newEmail string) {
	// Go automatically dereferences: user.Email is same as (*user).Email
	user.Email = newEmail
	// You can also explicitly dereference: (*user).Email = newEmail
}

// Pattern 2: Function that returns pointer to struct
// *User in return type = "return a pointer to User"
// This is common in constructors (like NewUser)
func NewUser(id int, name, email string) *User {
	// Create a struct value
	u := User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	// &u returns the ADDRESS of the struct (pointer)
	// This is what we return
	return &u
}

// Pattern 3: Method with pointer receiver
// (u *User) means "this method works on a pointer to User"
// Changes will affect the original struct
func (u *User) UpdateName(newName string) {
	u.Name = newName // Go auto-dereferences, same as (*u).Name
}

// Pattern 4: Method with value receiver (for comparison)
// (u User) means "this method works on a COPY of User"
// Changes won't affect the original
func (u User) GetName() string {
	return u.Name // Just reading, no pointer needed
}

func main() {
	fmt.Println("============ BASIC POINTERS ============")
	num := 1
	fmt.Println("memory address is", &num)

	// Passing by value (copy)
	changeNumbyvalue(num)
	fmt.Println("after change num by value", num, "<-- Still 1!")

	// Passing by reference (pointer)
	// &num passes the ADDRESS of num
	changeNumByReference(&num)
	fmt.Println("after change num by reference", num, "<-- Now 5!")

	fmt.Println("\n============ STRUCTS WITH POINTERS ============")

	// Creating a struct value (not a pointer)
	user1 := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	fmt.Println("user1:", user1)

	// Pattern 1: Pass pointer to function
	// &user1 gets the ADDRESS of user1
	updateUserEmail(&user1, "alice.new@example.com")
	fmt.Println("user1 after update:", user1)

	// Pattern 2: Create struct using constructor (returns pointer)
	// NewUser returns *User (pointer)
	user2 := NewUser(2, "Bob", "bob@example.com")
	// user2 is of type *User (pointer to User)
	fmt.Println("user2 (pointer):", user2)
	fmt.Println("user2 value:", *user2) // *user2 dereferences to get the value

	// Pattern 3: Call method with pointer receiver
	// Go automatically passes &user1 (address) when you call the method
	user1.UpdateName("Alice Updated")
	fmt.Println("user1 after method:", user1)

	// Pattern 4: Working with pointer variables
	var user3 *User // user3 is a pointer (holds address)
	user3 = NewUser(3, "Charlie", "charlie@example.com")
	fmt.Println("user3:", user3)
	fmt.Println("user3 value:", *user3)

	// You can also create pointer directly
	user4 := &User{ID: 4, Name: "Diana", Email: "diana@example.com"}
	fmt.Println("user4:", user4)

	// Modifying through pointer
	user4.Name = "Diana Updated"             // Go auto-dereferences
	(*user4).Email = "diana.new@example.com" // Explicit dereference (same thing)
	fmt.Println("user4 after update:", user4)

	fmt.Println("\n============ KEY TAKEAWAYS ============")
	fmt.Println("&variable  -> Get address (pointer)")
	fmt.Println("*pointer   -> Get value at address (dereference)")
	fmt.Println("*Type      -> Type declaration for pointer")
	fmt.Println("In structs: Use *StructName for parameters, returns, and receivers")
	fmt.Println("Go auto-dereferences: user.Name works even if user is *User")
}
