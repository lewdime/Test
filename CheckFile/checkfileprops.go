package checkfileprops

import (
	"fmt"
	"log"
	"os"
)

func main() {
	GenerateFileStatusReport("testfile.txt")
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

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
