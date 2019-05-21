package main

import "fmt"
import "errors"

func main() {
	var i, ch, n int
	var a [100]int
	for true {
		fmt.Println("Enter your choice 1:add 2:subract 3:division 4:multiplication 5:exit")
		fmt.Scanln(&ch)
		if ch != 5 {
			fmt.Println("Enter the number of elements on which the operation is to be performed")
			fmt.Scanln(&n)
			fmt.Println("Enter the elements")
			for i = 0; i < n; i++ {
				fmt.Scanln(&a[i])
			}
		} else {
			break
		}
		switch ch {
		case 1:
			fmt.Println("The result of addition is ", my_add(a, n))
		case 2:
			fmt.Println("The result of subraction is", my_subract(a, n))
		case 3:
			var e error
			var res int
			res, e = my_division(a, n)
			if e != nil {
				fmt.Println(e)
			} else {
				fmt.Println("The result of division is ", res)
			}
		case 4:
			fmt.Println("The result of multiplocation is", my_multiply(a, n))
		default:
			fmt.Println("Enter a valid choice")
		}
	}

}
func my_add(a [100]int, n int) int {
	var i int
	var sum int = 0
	for i = 0; i < n; i++ {
		sum = sum + a[i]
	}
	return sum
}
func my_subract(a [100]int, n int) int {
	var i int
	var diff int = a[0]
	for i = 1; i < n; i++ {
		diff = diff - a[i]
	}
	return diff
}
func my_division(a [100]int, n int) (int, error) {
	var i int
	var res int = a[0]
	for i = 0; i < n; i++ {
		if a[i] == 0 {
			var e error = errors.New("Divide by zero")
			return -1, e
		}
	}
	for i = 1; i < n; i++ {
		res = res / a[i]
	}
	return res, nil
}
func my_multiply(a [100]int, n int) int {
	var i int
	var res int = 1
	for i = 0; i < n; i++ {
		res = res * a[i]
	}
	return res
}
