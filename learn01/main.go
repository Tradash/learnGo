package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

// this is a comment

func main() {
	fmt.Println("Hello World")
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("Ошибка получения времени", err)
	} else {
		fmt.Println("Текущее время", time, err)
	}

}
