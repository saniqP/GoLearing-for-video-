package main

import "fmt"

/*
func check_null_list(list []int) bool {

	null := true

	for _, val := range list {
		if val != 0 {
			null = false
			break
		}
	}

	if null {
		fmt.Println("Null list")
		return true
	} else {
		return false
	}
}

func null_list(list []int, index int) int {

	if check_null_list(list) {
		return 0
	} else {
		if index+1 < len(list) {

			fmt.Println("------------------------------------------------")
			fmt.Printf("Index %v: %v >> %v\nIndex %v: %v >> %v\n", index, list[index], 0, index+1, list[index+1], list[index])

			list[index+1] = list[index]
			list[index] = 0
		} else {
			list[index] = 0
		}

		return null_list(list, index+1)
	}
}

func main() {
	var list []int = []int{1, 2, 3, 4, 5}
	null_list(list, 0)

	var x int = 10
	var y *int = &x
	fmt.Println(y)
	fmt.Println(*y)

}
*/

type person struct {
	name    string
	age     int
	freinds *[]person
}

func (p person) print_freinds_info() {
	fmt.Println(p.name, "freinds:")
	for _, freind := range *p.freinds {
		fmt.Printf("Freind name: %s\nFreind age: %v\n", freind.name, freind.age)
		if freind.freinds != nil {
			println("------------------")
			print(" ")
			freind.print_freinds_info()
			println("------------------")
		}
	}
}

func main() {
	var human1 person = person{"Ivan", 11, nil}
	var human2 person = person{"Sania", 12, &[]person{
		person{"Matvei", 12, nil},
		human1,
	}}
	var human3 person = person{"Dima", 12, &[]person{
		human1,
		human2,
	}}

	human3.print_freinds_info()
}
