// package name and folder name will be same for good practices
package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("logging user using name and password")
}

func tryingToUsethisfuncIntoAnotherPackage() {
	fmt.Println("trying to use this function which is having small letter to another packages")
}
