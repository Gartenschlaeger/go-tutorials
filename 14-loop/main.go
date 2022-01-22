package main

import "fmt"

func main() {

	days := []string{
		"Montag",
		"Dienstag",
		"Mittwoch",
		"Donnerstag",
		"Freitag",
		"Samstag",
		"Sonntag"}

	// for zählschleife

	for i := 0; i < len(days); i++ {
		fmt.Printf("days[%v] = %v\n", i, days[i])
	}

	// range foreach

	for index, day := range days {
		fmt.Printf("days[%v] = %v\n", index, day)
	}

	// endlosschleife

	counter := -1
	for {
		counter++
		if counter == 1 {
			continue
		}
		if counter > 3 {
			break
		}

		fmt.Printf("counter = %v\n", counter)
	}

	max := 1
	for max < 10 {
		fmt.Printf("max = %v\n", max)
		max++
	}

}
