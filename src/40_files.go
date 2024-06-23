package main

import (
	"fmt"
	"io/fs"
	"os"
)

func files() {
	createFile("sampleFile.txt", "Content file")
	writeBytes("sampleBytes.txt", []byte("Content bytes"))
	appendToFile("sampleAppend.txt", "Appending data")
}

func createFile(filename string, data string) int {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer file.Close()

	size, err := file.WriteString(data)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return size
}

func writeBytes(filename string, data []byte) int {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer file.Close()

	size, err := file.Write(data)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return size
}

func appendToFile(filename string, data string) int {
	if _, err := os.Stat(filename); err != nil {
		file, _ := os.Create(filename)
		defer file.Close()
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer file.Close()

	size, err := fmt.Fprintln(file, data)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return size
}
