package main

import "fmt"

const hash_len = 100

func hash(i int) int {
	return i % hash_len
}

func main() {
	l := []int{16, 29, 31, 74, 4, 14, 59, 58, 78}
	var table [15]int
	fmt.Printf("%2v\n", l)
	for _, v := range l {
		t := hash(v)
		if table[t] == 0 {
			table[t] = v
		} else {
			i := 1
			for {
				if table[hash(t+i)] == 0 {
					table[hash(t+i)] = v
					break
				}
				i++
			}
		}
		fmt.Printf("%2v\n", table)
	}
}
