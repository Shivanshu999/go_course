package main

import (
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("is directory:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file mode:", fileInfo.Mode())
	// fmt.Println("file modified:", fileInfo.ModTime())

	//read file
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("bytes read:", d, string(buf[i]))

	// }

	//how to read file using os.ReadFile

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// println(string(data))

	//read folder

	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// files, err := dir.Readdir(-1)

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	//create file

	f, err := os.Create("index.ts")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// f.WriteString("console.log('hello world')")

	// bytes := []byte("console.log('hello broos')")
	// f.Write(bytes)

}
