package main

import (
	"fmt"
)

func happy() {
	happy := []string{
		" 00000  0000   0       00000 00000 ",
		" 0      0      0  	 0      0        ",
		" 00000  0000   0         0     0   ",
		" 0      0      0    0    0    0    ",
		" 0      0000   000000  00000 00000 ",
	}

	for i := range happy {
		fmt.Println(happy[i])
	}
}

func day() {
	day := []string{
		"0000   00000   000  ",
		"0   0    0    0   0 ",
		"0    0   0    00000 ",
		"0   0    0    0   0 ",
		"0000   00000  0   0 ",
	}

	for i := range day {
		fmt.Println(day[i])
	}
}

func to() {
	to := []string{
		"0000      00  ",
		"0   0   0    0",
		"0    0  0    0",
		"0   0   0    0",
		"0000	  00  ",
	}

	for i := range to {
		fmt.Println(to[i])
	}
}

func work() {
	work := []string{
		"0000000  0000      000   00000     000    0        0    0    00000   ",
		"   0     0   0    0   0  0    0   0   0   0        0    0   0     0  ",
		"   0     0000     00000  00000    00000   0        000000   0     0  ",
		"   0     0   0    0   0  0    0   0   0   0    0   0    0   0     0  ",
		"   0     0    0   0   0  00000    0   0   000000   0    0    00000   ",
	}

	for i := range work {
		fmt.Println(work[i])
	}
}

func dayWork() {
	fmt.Println("=======================================")
	fmt.Println()
	happy()
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println()
	day()
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println()
	to()
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println()
	work()
	fmt.Println()
	fmt.Println("=======================================")
}

func main() {
	dayWork()
}
