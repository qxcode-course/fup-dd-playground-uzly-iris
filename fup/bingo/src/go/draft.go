package main
import "fmt"
func main() {
    bingo := [4][4] int {
		{1, 9, 27, 23},
		{34, 20, 37, 47},
		{30, 87, 55, 69},
		{13, 60, 99, 66},
	}

	var chutes[6] int
	nums := 0
	c := 0

	for i := range chutes {
		fmt.Scan(&nums)
		chutes[i] = nums
	}

	for _, e := range chutes {
		for i := 0 ; i < 4 ; i++ {
			for j := 0 ; j < 4 ; j++ {
				if bingo[i][j] == e {
					c++
				}
			}
		}
	}

	fmt.Println(c)
}