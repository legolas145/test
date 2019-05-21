package main

import "fmt"

func main() {
	var umap map[int]string
	umap = make(map[int]string)
	var i int
	umap[1] = "one"
	umap[2] = "two"
	fmt.Println(umap[1], umap[2])
	for i = range umap {
		fmt.Println(i, umap[i])
	}

}
