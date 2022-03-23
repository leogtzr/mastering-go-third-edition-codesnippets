/* Usage example:

$ go run dates.go
Usage: dates parse_string
$ go run dates.go 14:10
Full: 0000-01-01 14:10:00 +0000 UTC
Time: 14 10
Epoch time: 1607964956
Date: 14 December 2020
Time: 18:55
Execution time: 163.032µs
$ go run dates.go "14 December 2020"
Full: 2020-12-14 00:00:00 +0000 UTC
Time: 14 December 2020
Epoch time: 1607964985
Date: 14 December 2020
Time: 18:56
Execution time: 180.029µs

*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: dates parse_string")
		os.Exit(0)
	}

	dateString := os.Args[1]
	// Is this a date only?
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	// Is this a date + time
	d, err = time.Parse("02 January	2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is this a date + time with month represented as a number?
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is it time only?
	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// Convert Epoch time to time.Time
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
	duration := time.Since(start)
	fmt.Println("Execution time:", duration)

    now := time.Now()
    loc, _ := time.LoadLocation("America/New_York")
    fmt.Printf("New York Time: %s\n", now.In(loc))
}
