package array

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	var arr = []int{1, 2, 3}
	fmt.Println("size 0 -> ", Chunk[int](arr, 0))
	fmt.Println("size 1 -> ", Chunk[int](arr, 1))
	fmt.Println("size 2 -> ", Chunk[int](arr, 2))
	fmt.Println("size 3 -> ", Chunk[int](arr, 3))
	fmt.Println("size 4 -> ", Chunk[int](arr, 4))
}

func TestIntJoin(t *testing.T) {
	var arr = []int{0, 1, 2, 3, -2}
	fmt.Println("size 0 -> ", IntJoin[int](arr, ""))
	fmt.Println("size 1 -> ", IntJoin[int](arr, " "))
	fmt.Println("size 2 -> ", IntJoin[int](arr, "|"))
	fmt.Println("size 3 -> ", IntJoin[int](arr, "-"))
	fmt.Println("size 4 -> ", IntJoin[int](arr, ","))
	fmt.Println("size 5 -> ", IntJoin[int](arr, ", "))
	var arr2 []int
	fmt.Println("size 6 -> ", IntJoin[int](arr2, " "))
	fmt.Println("size 7 -> ", IntJoin[int](arr2, "-"))
}
