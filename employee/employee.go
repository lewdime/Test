package employee

import (
	"fmt"
)

//  the stuct below cannot be exported and consumed by importing

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

/*
   the func below let's the employee struc be created from a func routine call
   and returns an "employee" struct type to be consumed by the caller func
*/

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

//  the func below will be a method in the emloyee struct to output a record
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
