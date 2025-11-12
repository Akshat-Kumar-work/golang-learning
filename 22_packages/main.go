package main

import "golang-learning/22_packages/auth"

//to start with packages
// first create a mod for project
// go mod init project-name / github repo url

//NOTE : if package is same we can use any functions or file

//IMP : if we write function name with small first letter that function is scoped to that particular package
// if want to to use that function into other package write function first letter in CAPS

//to install any 3rd party package from go registry go get package url

func main() {
	auth.LoginWithCredentials("akshat", "ok")
	//auth.tryingToUsethisfuncIntoAnotherPackage()
}
