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
