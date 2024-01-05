package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5}
	res := make([]*int, 0)
	for _, v := range test {
		fmt.Printf("append %d, addr:%p\n", v, &v)
		res = append(res, &v)
	}

	for _, v := range res {
		fmt.Println(*v)
	}
}
