package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	fibs = append(fibs, (fibs[3] + fibs[4]))

	for range 3 {
		length := len(fibs)
		nextNum := fibs[length-1] + fibs[length-2]

		fibs = append(fibs, nextNum)
	}
	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
