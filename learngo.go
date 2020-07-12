package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

//	create calculate method in Fixedbilling struct that returns biddedAmount field
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

//	create source method in Fixedbilling struct that returns source of income in projectName field
func (fb FixedBilling) source() string {
	return fb.projectName
}

//	created calculate method in TimeAndMaterial struct that calculate the and returns income
func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

//	created source method in TimeAndMaterial struct that returns source of income in projectName field
func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// create a function that calculates and returns the total income from slices of Income parameter argument
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d\n", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	//	manually creating series of records of Income interfaces in from different struct type records
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
