package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func CountMax(weights []int, values []int, Capacity int, n int) int {
	arr := make([][]int, Capacity+1)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for w := 1; w < Capacity+1; w++ {
		for k := 0; k < n; k++ {
			if weights[k] <= w {
				if k == 0 {
					arr[w][k] = values[k]
				} else {
					arr[w][k] = max(arr[w][k-1], (arr[w-weights[k]][k-1])+values[k])
				}
			}
		}

	}
	return arr[Capacity][n-1]
}

func main() {
	var n int
	var values []int
	var weights []int
	fmt.Println("Введите количество вещей: ")
	fmt.Scan(&n)
	fmt.Println("Введите вес и стоимость вещей ( _вес_ , _стоимость_ ): ")
	for i := 0; i < n; i++ {
		var weigth, value int
		fmt.Print(i + 1)
		fmt.Print(": ")
		fmt.Scan(&weigth, &value)
		weights = append(weights, weigth)
		values = append(values, value)
	}
	var MaxCapacity int
	fmt.Println("Введите вес рюкзака: ")
	fmt.Scan(&MaxCapacity)
	fmt.Println("Максимальная ценность вещей в рюкзаке: ")
	fmt.Println(CountMax(weights, values, MaxCapacity, n))
}

/*
Введите количество вещей:
5
Введите вес и стоимость вещей ( _вес_ , _стоимость_ ):
1: 6 5
2: 4 3
3: 3 1
4: 2 3
5: 5 6
Введите вес рюкзака:
15

*/

/*
Введите количество вещей:
3
Введите вес и стоимость вещей ( _вес_ , _стоимость_ ):
1: 4 4000
2: 1 2500
3: 3 2000
Введите вес рюкзака:
4

*/
