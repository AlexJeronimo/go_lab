package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ReportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}
func ScanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// return err
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
	//return nil
}

func main() {
	/*
		err := ScanDirectory("C:\\Users\\sysadmin\\go\\src")
		if err != nil {
			log.Fatal(err)
		}
	*/
	defer ReportPanic()
	//panic("some other issue")
	ScanDirectory("C:\\Users\\sysadmin\\go\\src")
}
