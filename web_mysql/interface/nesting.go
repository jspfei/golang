package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Writer()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f *File) Read() {
	fmt.Println("文件读取")
}

func (f *File) Writer() {
	fmt.Println("文件写入")
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Writer()
}

func main() {
	var f File
	Test(&f)
}
