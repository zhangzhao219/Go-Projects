package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// reader := bufio.NewReader(f)
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(str)
	// }
	// f.Close()

	// fileName := "temp.txt"
	// file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	// writer := bufio.NewWriter(file)
	// reader = bufio.NewReader(file)
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(str)
	// }
	// for i := 0; i < 5; i++ {
	// 	writer.WriteString("hellonnn4512\n")
	// }
	// writer.Flush()
	// file.Close()

	// file1 := "test.txt"
	// file2 := "temp.txt"
	// read, _ := ioutil.ReadFile(file1)
	// err := ioutil.WriteFile(file2, read, 0644)
	// fmt.Println(err)
	// _, err = os.Stat("tem.txt")
	// if os.IsNotExist(err) {
	// 	fmt.Println("fdsdsff")
	// }
	// fmt.Println(os.Args)
	// var user string
	// var pwd string
	// var host int
	// flag.StringVar(&user, "u", "", "user")
	// flag.StringVar(&pwd, "p", "", "pwd")
	// flag.IntVar(&host, "h", 0, "host")

	// flag.Parse()

	// fmt.Println(user, pwd, host)

	copyFile("temp2.txt", "temp.txt")

}

func copyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("error")
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error2")
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}
