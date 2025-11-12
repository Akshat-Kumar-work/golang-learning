package main

import (
	"fmt"
	"os"
)

// NOTE: In go whenver the err comes we usually don't throw we always return

func main() {
	f, err := os.Open("./15_enum/main.go")
	if err != nil {
		//log the err
		//panic
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("file name", fileInfo.Name())
}
