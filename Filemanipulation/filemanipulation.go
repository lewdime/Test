package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	//open a file for read only
	fh1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer fh1.Close()

	//Create a new file
	fh2, err := os.Create("test2.txt")
	PrintFatalError(err)
	defer fh2.Close()

	//open file fro read write

	fh3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)

	// os.O_RDONLY		// Read Only
	// os.O_WRONLY		// Write Only
	// os.O_RDWR		// Read Write
	// os.O_APPEND		// Append to end of file
	// os.O_CREATE		// Create if none exists
	// os.O_TRUNC		// Truncate file when opening

	// 0666 => Owner: (read &write), Group: (read & write), and other (read & write)

	PrintFatalError(err)
	defer fh3.Close()

	//rename a file
	//err = os.Rename("test1.txt", "test1New.txt")
	//PrintFatalError(err)

	//move a file
	err = os.Rename("./test1.txt", "./testfolder/test1.txt")
	PrintFatalError(err)

	//copy a file
	CopyFile("test3.txt", "./testfolder/test3.txt")

	//delte a file
	err = os.Remove("test2.txt")
	PrintFatalError(err)

	bytes, err := ioutil.ReadFile("test3.txt")
	fmt.Println(string(bytes))

	scanner := bufio.NewScanner(fh3)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println("Found line:", count, scanner.Text())
	}

	//buffered write, efficient store in memory, saves disk I/O
	writebuffer := bufio.NewWriter(fh3)
	for i := 1; i <= 5; i++ {
		writebuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writebuffer.Flush()

	GenerateFileStatusReport("test3.txt")

	filestat1, err := os.Stat("test3.txt")
	PrintFatalError(err)
	for {
		time.Sleep(1 * time.Second)
		filestat2, err := os.Stat("test3.txt")
		PrintFatalError(err)
	}

}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}

//Copy file fname1 to fname2
func CopyFile(fname1, fname2 string) {
	fOld, err := os.Open(fname2)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	//copy bytes from source to destination

	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	//flush file contents to desc
	err = fNew.Sync()
	PrintFatalError(err)

}

func GenerateFileStatusReport(fname string) {
	// Stat returns file info. It will return
	// and error if there is no file
	filestats, err := os.Stat(fname)
	PrintFatalError(err)

	fmt.Println("What's the file name?", filestats.Name())
	fmt.Println("Am i a directory?", filestats.IsDir())
	fmt.Println("What are the permissions?", filestats.Mode())
	fmt.Println("When was the last time the file was modified?", filestats.ModTime())
}
