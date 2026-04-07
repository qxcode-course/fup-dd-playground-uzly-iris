package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	contagem(n)

}

func contagem(n int){
		fmt.Print("[ ")
	for i := 0; i < 10; i++ {
		if i != n {
			fmt.Print(i, " ")
		}

	}
	if n != 10 {
		fmt.Print("ceu ")
	}
	fmt.Println("]")
}
