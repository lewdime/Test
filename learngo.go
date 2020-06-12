package main

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
