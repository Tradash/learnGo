package main

import (
	"fmt"
	"sort"
	"strings"
)

const str = "ssss fff ggg fff ssss  ggg ggg ggg ert ert dfsd ssss ggg ssss fff"

type Data struct {
	word  string
	count int
}

func counter(s string, numberTop int) []Data {
	rez := map[string]int{}
	mystr := strings.Split(s, " ")
	for i := 0; i < len(mystr); i++ {
		if mystr[i] != "" && mystr[i] != " " {
			rez[mystr[i]] = rez[mystr[i]] + 1
		}
	}
	var arr []Data
	for key, val := range rez {
		var d Data
		d.word = key
		d.count = val
		arr = append(arr, d)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})
	if numberTop > len(arr) {
		numberTop = len(arr)
	}
	return arr[0:numberTop]
}

func main() {
	println(str)
	itog := counter(str, 2)
	fmt.Printf("%v", itog)
}
