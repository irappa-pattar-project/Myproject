package Hello

import "fmt"

func Helloworld() {
	total := hello()
	fmt.Println("Total", total)
	fmt.Println("end of statement")
	//fmt.Println("Total", total)
}

func hello() int {
	total := 0
	for i := 0; i < 5; i++ {
		total += i
	}
	fmt.Println("total", total)
	return total
}
