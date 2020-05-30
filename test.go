/* 

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(a)) //type and size of a
	fmt.Printf("type of b is %T, size of b is %d\n", b, unsafe.Sizeof(b)) //type and size of b
}




package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}



package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

*/

/*

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}



package main

import "fmt"

func reverse(x int) {

	fmt.Println("counting")

	for i := 0; i < x; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	reverse(4)

}




package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

*/

/*

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

*/

/*

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

*/

/*

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

*/

/*

package main // Slices

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // take 2nd until 3rd
	fmt.Println(s)
}



package main // Slices are like reference to Arrays

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // The value Paul in array 1 in b is changed to XXX
	fmt.Println(a, b)
	fmt.Println(names)
}



package main //Slice Literals is like an array without lenght

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
	i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}





package main //event using channel

import (
	"log"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "", log.Ltime|log.Lmicroseconds)
	l.Println("program start")
	event := make(chan int)
	go func() {
		l.Println("task start")
		<-event
		l.Println("event reset by task")
	}()
	l.Println("program sleeping")
	time.Sleep(1 * time.Second)
	l.Println("program signaling event")
	event <- 0
	time.Sleep(100 * time.Millisecond)
}





package main 

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	var x, y float64
	fmt.Print("Enter 2 numbers: ")
	fmt.Scanln(&x, &y)
	return fn(x, y)
}

func main() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("The Squreroot is: ", hypot(5, 12))
	fmt.Println("The Squreroot of the sum of twice their value is: ", compute(hypot))
	fmt.Println(compute(math.Pow))
}





package main // Reading Console Input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

}





package main // Executing System Commands With Golang

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("ls", "-ltr").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	// let's try the pwd command herer
	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output = string(out[:])
	fmt.Println(output)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}



package main	// Parsing JSON Files

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

// Users struct which contains
// an array of users
type Users struct {
    Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var users Users

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &users)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

}



package main	// Go init Function https://tutorialedge.net/golang/the-go-init-function/

import "fmt"

// this variable is initialized first due to
// order of declaration
var initCounter int

func init() {
    fmt.Println("Called First in Order of Declaration")
    initCounter++
}

func init() {
    fmt.Println("Called second in order of declaration")
    initCounter++
}

func main() {
    fmt.Println("Does nothing of any significance")
    fmt.Printf("Init Counter: %d\n", initCounter)
}


package main // Channel

import "fmt"

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}


package main  // Sorting With the Sort Package

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Age int 
} 

type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
} 

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {
    programmers := []Programmer{
		Programmer{Age: 30,},
		Programmer{Age: 20,},
		Programmer{Age: 50,},
		Programmer{Age: 1000,},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)
}