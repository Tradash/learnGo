package main

import (
	"fmt"
	"strconv"
)

// const escp =
func escUnpack(s string) string {
	rez := ""
	newEsc := 0
	symbol := ""
	for i := 0; i < len(s); i++ {
		if d := "\\"; s[i] == d[0] && newEsc != 1 {
			newEsc = 1
			continue
		}
		if newEsc == 1 {
			symbol = string(s[i])
			newEsc = 2
			rez = rez + string(s[i])
			continue
		}
		if newEsc == 2 {
			numb, err := strconv.ParseInt(string(s[i]), 10, 32)
			if err == nil {
				for j := 0; j < int(numb)-1; j++ {
					rez = rez + symbol
				}
				newEsc = 0
				continue
			}

		}
		rez = rez + string(s[i])
	}
	return rez
}

func unpack(s string) string {
	rez := ""
	startNewPack := 1
	for i := 0; i < len(s); i++ {

		numb, err := strconv.ParseInt(string(s[i]), 10, 32)
		if err == nil {
			startNewPack = int(numb)
			continue
		}
		for j := 0; j < startNewPack; j++ {
			rez = rez + string(s[i])
		}
		startNewPack = 1
	}
	return rez
}

func main() {
	test1 := "a4bc2d5e"
	test2 := "abcd"
	test3 := "45"
	fmt.Println("Поехали")
	fmt.Println(test1, unpack(test1))
	fmt.Println(test2, unpack(test2))
	fmt.Println(test3, unpack(test3))
	fmt.Println(`qwe\4\5`, escUnpack(`qwe\4\5`))
	fmt.Println(`qwe\45`, escUnpack(`qwe\45`))
	fmt.Println(`qwe\45\-4\5`, escUnpack(`qwe\45\-4\5`))
	fmt.Println(`qwe\\5\73`, escUnpack(`qwe\\5\73`))
}
