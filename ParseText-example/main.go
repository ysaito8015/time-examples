package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now())
	onceUponATime := time.Date(2009, 11, 17, 20, 34, 58, 561387237, time.UTC)
	fmt.Println(onceUponATime)

	atTime := "2023-06-02 13:23:08"
	fmt.Println(atTime)
	date := strings.Split(atTime, " ")[0]
	time := strings.Split(atTime, " ")[1]
	fmt.Println(date)
	date2 := strings.Split(date, "-")
	fmt.Println(date2[0])
	fmt.Println(strings.Join(date2, "/"))
	fmt.Println(time)
}
