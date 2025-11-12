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
	fmt.Println("is directory", fileInfo.IsDir())
	fmt.Println("file permission", fileInfo.Mode())

	//close file, when the function return anything
	defer f.Close()

	//read file content
	//we read file content and keep it in buffer (memory)
	buf := make([]byte, 10)                      //10 bytes slice , or we can use fileInfoSize to read whole file
	numberOfByteActuallyRead, err := f.Read(buf) //read file 10 byte data only and keep it in buf slice
	if err != nil {
		panic(err)
	}
	//here in range we are saying from 0 to number of actually bytes as it contain garbage values too
	for i := range buf[:numberOfByteActuallyRead] {
		fmt.Println("data", i, string(buf[:numberOfByteActuallyRead]), numberOfByteActuallyRead)

	}

	//we also have another simple high level method to read file
	// this func has drawback as it loads whole file at once in memory
	//so we can go with this if file size is small
	data, err := os.ReadFile("./11_closures/main.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
