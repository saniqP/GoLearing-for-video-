package main

import (
	"fmt"
)

func main() {

	//Прошлый десяток
	var past_des int = 0

	//Двухмерный масив
	var list [7][3]int = [7][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
		{19, 20, 21},
	}

	//Перебор массивов в массиве
	for i := 0; i < len(list); i++ {

		fmt.Println("Array", i+1, ":")

		//если проошлый десяток + 10 больше или равен то мы пишем какой десяток
		if list[i][2] >= past_des+10 {
			fmt.Printf("This is %v десяток\n", list[i][2]/10)
			past_des = past_des + 10
		}

		//перебераем элементы в текущем массиве
		for x := 0; x < len(list[i-1]); x++ {

			fmt.Println("Element", x, ":", list[i][x])
		}
	}

}
