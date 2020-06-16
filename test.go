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



package main // Struct

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





package main // Arrays

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


package main  // Sorting With Sort Package

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




package main // Using Lock and Unlock Methon in sync.Mutex

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	i int
	sync.Mutex
}

func main() {

	sc := new(safeCounter)

	for i := 0; i < 100; i++ {
		go sc.Increment()
		go sc.Decrement()
	}

	fmt.Println(sc.GetValue())
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	sc.i++
	sc.Unlock()
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func (sc *safeCounter) GetValue() int {
	sc.Lock()
	v := sc.i
	sc.Unlock()
	return v
}




package main  // Using Once methon in the Sync Package

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}




package main // Multiple Read and a Write using sync.RWMutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

func main() {
	mc := MapCounter{m: make(map[int]int)}

	// Referencing to the address of mc by & as the parameter of the func
	// is using a pointer * to get the value of the same object
	go runWriters(&mc, 10)

	// Running two instances of go routine running as the same time
	go runReaders(&mc, 10)
	go runReaders(&mc, 10)
	time.Sleep(15 * time.Second)
}

// Creating the functions that will be used in the main function
func runWriters(mc *MapCounter, n int) {
	for i := 0; i < n; i++ {
		mc.Lock()
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func runReaders(mc *MapCounter, n int) {
	for {
		mc.RLock()
		//randomly choosing an item in the map  counter
		v := mc.m[rand.Intn(n)]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}



package main  // Using Wait Group method of sync

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {

		// Incrementing the WaitGroup counter
		wg.Add(1)

		// Launch a go routine
		go func(i int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Do some work
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for ", i)

		}(i)
	}

	wg.Wait()
}



package main	//  Ticker method using time package

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tock")
	}
}

func main() {
	fmt.Println("Go Tickers Tutorial")

	go backgroundTask()
}

	// This print statement will be executed before
	// the first `tock` prints in the console
	fmt.Println("The rest of my application can continue")
	// here we use an empty select{} in order to keep
	// our main function alive indefinitely as it would
	// complete before our backgroundTask has a chance
	// to execute if we didn't.
	select {}



package main	// Using Maps with interface

import "fmt"

type Service interface {
	SayHi()
}

type MyService struct{}

func (s MyService) SayHi() {
	fmt.Println("Hi")
}

type SecondService struct{}

func (s SecondService) SayHi() {
	fmt.Println("Hello From the 2nd Service")
}

func main() {
	fmt.Println("Go Maps Tutorial")
	// we can define a map of string uuids to
	// the interface type 'Service'
	interfaceMap := make(map[string]Service)

	// we can then populate our map with
	// simple ids to particular services
	interfaceMap["SERVICE-ID-1"] = MyService{}
	interfaceMap["SERVICE-ID-2"] = SecondService{}

	// Incoming HTTP Request wants service 2
	// we can use the incoming uuid to lookup the required
	// service and call it's SayHi() method
	// interfaceMap["SERVICE-ID-2"].SayHi()

	// Below we assign the "key" as the string service name of the map
	// and the "service" as the struct value
	// We could iterate over all the interfaces within our map
	// and call every SayHi() method
	for key, service := range interfaceMap {
		fmt.Println(key)
		service.SayHi()
	}

}




package main	// Go Routines

import (
	"fmt"
	"time"
)

// a very simple function that we'll
// make asynchronous later on
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// notice how we've added the 'go' keyword
	// in front of both our compute function calls
	go compute(10)
	go compute(10)

	// we can make our anonymous function concurrent using go as below
	// go func() {
	//	fmt.Println("Executing my Concurrent anonymous function")
	// }()

	// we scan fmt for input and print that to our console
	// this is so that our program waits for keyboard imput
	// before it kills our poor go routines
	var input string
	fmt.Scanln(&input)

}



package main	// Go Routines with Waitgroup

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var urls = []string{
    "https://google.com",
    "https://tutorialedge.net",
    "https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    wg.Done()
    fmt.Println(resp.Status)
    return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HomePage Endpoint Hit")
    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1)
        go fetch(url, &wg)
    }

    wg.Wait()
    fmt.Println("Returning Response")
    fmt.Fprintf(w, "Responses")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}