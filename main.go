package main

import (
	"errors"
	"fmt"
)

func main() {
	a, err := triangle(3, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

// triangle 三角形の面積を返す
func triangle(bottom, height int) (int, error) {
	if bottom == 0 {
		return 0, errors.New("底辺の値が0です")
	}

	if height == 0 {
		return 0, errors.New("高さの値が0です")
	}

	return bottom * height / 2, nil
}
